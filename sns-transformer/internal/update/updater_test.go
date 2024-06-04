package update

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
	coresnsv1client "github.com/ironzhang/sns/kernel/clients/coresnsclients/clientset/versioned/typed/core.sns.io/v1"
)

func NewTestUpdater(t *testing.T) *Updater {
	cfg, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		t.Fatalf("build config from flags: %v", err)
	}

	snscorev1Client, err := coresnsv1client.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("new for config: %v", err)
	}

	return NewUpdater(snscorev1Client)
}

func TestUpdater(t *testing.T) {
	c := coresnsv1.SNSCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hna-v.sns.grpc.test",
			Namespace: "sns",
			Labels: map[string]string{
				"cluster": "hna-v",
				"domain":  "sns.grpc.test",
			},
		},
		Spec: coresnsv1.ClusterSpec{
			//Tags: map[string]string{
			//	"Environment": "product",
			//	"Lidc":        "hna",
			//},
			Endpoints: []coresnsv1.Endpoint{
				{
					Addr:   "127.0.0.1:8001",
					State:  coresnsv1.Enabled,
					Weight: 100,
					Tags:   map[string]string{"hostname": "localhost"},
				},
			},
		},
	}

	u := NewTestUpdater(t)
	err := u.UpdateCluster(context.Background(), c)
	if err != nil {
		t.Fatalf("update cluster: %v", err)
	}
}
