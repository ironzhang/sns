package app

import (
	"context"

	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/ironzhang/tlog"

	coresnsv1client "github.com/ironzhang/sns/kernel/clients/coresnsclients/clientset/versioned/typed/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/sns-transformer/internal/transform"
	"github.com/ironzhang/sns/sns-transformer/internal/update"
)

type Application struct {
	transformer *transform.Transformer
}

func buildRestClientConfig() (*restclient.Config, error) {
	cfg, err := restclient.InClusterConfig()
	if err == nil {
		tlog.Info("build rest client config in cluster")
		return cfg, nil
	}

	cfg, err = clientcmd.BuildConfigFromFlags("", Conf.Kube.KubeConfigPath)
	if err == nil {
		tlog.Info("build rest client config out of cluster")
		return cfg, nil
	}

	return nil, err
}

func newTransformer(cfg *restclient.Config) (*transform.Transformer, error) {
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	coresnsclient, err := coresnsv1client.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return transform.NewTransformer(Conf.Transform,
		k8sclient.NewWatchClient(clientset.CoreV1().RESTClient()),
		update.NewUpdater(coresnsclient)), nil
}

func (p *Application) Init() error {
	cfg, err := buildRestClientConfig()
	if err != nil {
		tlog.Errorw("build rest client config", "error", err)
		return err
	}

	p.transformer, err = newTransformer(cfg)
	if err != nil {
		tlog.Errorw("new transformer", "error", err)
		return err
	}

	return nil
}

func (p *Application) Fini() error {
	return nil
}

func (p *Application) Run(ctx context.Context) error {
	p.transformer.Start(ctx)

	<-ctx.Done()
	return nil
}
