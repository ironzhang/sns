package engine

import (
	"context"
	"os"
	"testing"

	"k8s.io/client-go/tools/clientcmd"

	coresnsv1client "github.com/ironzhang/sns/kernel/clients/coresnsclients/clientset/versioned/typed/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/filewrite"
	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/sns-agent/internal/paths"
)

func NewTestEngine(t testing.TB) *Engine {
	cfg, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		t.Fatalf("build config from flags: %v", err)
	}

	snscorev1Client, err := coresnsv1client.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("new for config: %v", err)
	}

	opts := Options{
		Namespace: "sns",
	}
	wc := k8sclient.NewWatchClient(snscorev1Client.RESTClient())
	pm := paths.NewPathManager("./testdata")
	fw := filewrite.NewFileWriter(pm.TemporaryPath())
	return NewEngine(opts, wc, pm, fw)
}

func TestEngine(t *testing.T) {
	ctx := context.Background()

	e := NewTestEngine(t)
	err := e.WatchDomain(ctx, "sns.http.nginx")
	if err != nil {
		t.Fatalf("watch domains: %v", err)
	}

	<-ctx.Done()
}

func TestMain(m *testing.M) {
	os.RemoveAll("./testdata")
	m.Run()
	//os.RemoveAll("./testdata")
}
