package engine

import (
	"context"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/filewrite"
	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/sns-agent/internal/paths"
)

type Options struct {
	Namespace   string
	DefaultZone string
	DefaultLane string
	DefaultKind string
}

type Engine struct {
	opts     Options
	wc       *k8sclient.WatchClient
	cw       clusterWatcher
	indexers map[string]cache.Indexer
}

func NewEngine(opts Options, wc *k8sclient.WatchClient, pm *paths.PathManager, fw *filewrite.FileWriter) *Engine {
	return &Engine{
		opts: opts,
		wc:   wc,
		cw: clusterWatcher{
			sorter: &clusterSorter{
				DefaultZone: opts.DefaultZone,
				DefaultLane: opts.DefaultLane,
				DefaultKind: opts.DefaultKind,
			},
			pathmgr: pm,
			fwriter: fw,
		},
		indexers: make(map[string]cache.Indexer),
	}
}

func (p *Engine) watchClusters(ctx context.Context, domain string) error {
	ls, err := newDomainLabelSelector(domain)
	if err != nil {
		return err
	}

	key := clusterIndexerKey(domain)
	p.indexers[key] = p.wc.Watch(ctx, p.opts.Namespace, "snsclusters", &coresnsv1.SNSCluster{}, ls, fields.Everything(), cache.Indexers{}, &p.cw)

	return nil
}

func (p *Engine) WatchDomain(ctx context.Context, domain string) (err error) {
	if err = p.watchClusters(ctx, domain); err != nil {
		return err
	}
	return nil
}

// RefreshClusters refresh the given domain's cluster file.
func (p *Engine) RefreshClusters(ctx context.Context, domain string) {
	key := clusterIndexerKey(domain)
	indexer, ok := p.indexers[key]
	if ok {
		p.cw.OnRefresh(indexer)
	}
}
