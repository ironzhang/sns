package transform

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	"github.com/ironzhang/tlog"

	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/pkg/snsutil"
	"github.com/ironzhang/sns/sns-transformer/internal/update"
)

type Options struct {
	SourceNamespace string
	TargetNamespace string
	DefaultZone     string
	DefaultLane     string
}

type Transformer struct {
	opts Options
	wc   *k8sclient.WatchClient
	pw   podWatcher
}

func NewTransformer(opts Options, wc *k8sclient.WatchClient, u *update.Updater) *Transformer {
	return &Transformer{
		opts: opts,
		wc:   wc,
		pw: podWatcher{
			targetNamespace: opts.TargetNamespace,
			cmnb: &snsutil.ClusterMetaNameBuilder{
				DefaultZone: opts.DefaultZone,
				DefaultLane: opts.DefaultLane,
				DefaultKind: snsutil.K8S_ClusterKind,
			},
			updater: u,
		},
	}
}

func (p *Transformer) Start(ctx context.Context) {
	tlog.WithContext(ctx).Debugw("transformer start")
	indexers := cache.Indexers{
		"cluster_index": clusterIndexFunc,
	}
	p.wc.Watch(ctx, p.opts.SourceNamespace, "pods", &corev1.Pod{}, labels.Everything(), fields.Everything(), indexers, &p.pw)
}
