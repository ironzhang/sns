package app

import (
	"context"
	"net"

	"github.com/ironzhang/superlib/httputils/echoutil"
	"github.com/ironzhang/tlog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	coresnsv1client "github.com/ironzhang/sns/kernel/clients/coresnsclients/clientset/versioned/typed/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/filewrite"
	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/sns-agent/internal/agent"
	"github.com/ironzhang/sns/sns-agent/internal/engine"
	"github.com/ironzhang/sns/sns-agent/internal/handlers/agentapi"
	"github.com/ironzhang/sns/sns-agent/internal/paths"
	"github.com/ironzhang/sns/sns-agent/internal/ready"
	"github.com/ironzhang/sns/sns-agent/internal/watch"
)

type Application struct {
	echosvr *echo.Echo
}

func newEngine(cfg *restclient.Config, pathmgr *paths.PathManager) (*engine.Engine, error) {
	coresnsclient, err := coresnsv1client.NewForConfig(cfg)
	if err != nil {
		tlog.Errorw("new core sns client", "error", err)
		return nil, err
	}

	return engine.NewEngine(Conf.Engine,
		k8sclient.NewWatchClient(coresnsclient.RESTClient()),
		pathmgr,
		filewrite.NewFileWriter(pathmgr.TemporaryPath()),
	), nil
}

func newAgent() (*agent.Agent, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", Conf.Kube.KubeConfigPath)
	if err != nil {
		tlog.Errorw("build config from flags", "kube_config_path", Conf.Kube.KubeConfigPath, "error", err)
		return nil, err
	}

	pathmgr := paths.NewPathManager(Conf.ResourcePath)
	eng, err := newEngine(cfg, pathmgr)
	if err != nil {
		tlog.Errorw("new engine", "error", err)
		return nil, err
	}

	inspection := ready.NewInspection(pathmgr)
	watcher := watch.NewWatcher(eng, inspection)
	return agent.New(watcher), nil
}

func (p *Application) Init() error {
	agent, err := newAgent()
	if err != nil {
		tlog.Errorw("new agent", "error", err)
		return err
	}

	ln, err := net.Listen("tcp", Conf.Listen.Addr)
	if err != nil {
		tlog.Errorw("net listen tcp", "addr", Conf.Listen.Addr, "error", err)
		return err
	}

	p.echosvr = echo.New()
	p.echosvr.Listener = ln
	p.echosvr.HTTPErrorHandler = echoutil.HTTPErrorHandler
	p.echosvr.Use(echoutil.Recover())
	p.echosvr.Use(echoutil.ServeMuxMiddleware(nil))
	p.echosvr.Use(middleware.RequestID())
	p.echosvr.Use(echoutil.TraceMiddleware())
	p.echosvr.Use(echoutil.AccessLogMiddleware())

	agentapi.Register(p.echosvr, agentapi.NewHandler(agent))

	return nil
}

func (p *Application) Fini() error {
	return nil
}

func (p *Application) RunHTTPServer(ctx context.Context) error {
	tlog.Infof("serve http server on %s", Conf.Listen.Addr)
	go p.echosvr.StartServer(p.echosvr.Server)

	<-ctx.Done()
	p.echosvr.Close()
	return nil
}
