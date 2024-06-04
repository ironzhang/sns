package transform

import (
	"context"
	"testing"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	coresnsv1client "github.com/ironzhang/sns/kernel/clients/coresnsclients/clientset/versioned/typed/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/sns-transformer/internal/update"
)

func NewTestTransformer(t *testing.T) *Transformer {
	cfg, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		t.Fatalf("build config from flags: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("new for config: %v", err)
	}

	snscorev1Client, err := coresnsv1client.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("new for config: %v", err)
	}

	opts := Options{
		SourceNamespace: "default",
		TargetNamespace: "sns",
	}
	wc := k8sclient.NewWatchClient(clientset.CoreV1().RESTClient())
	u := update.NewUpdater(snscorev1Client)
	return NewTransformer(opts, wc, u)
}

func TestTransformer(t *testing.T) {
	ctx := context.Background()

	c := NewTestTransformer(t)
	c.Start(ctx)

	time.Sleep(1 * time.Second)
	<-ctx.Done()
}
