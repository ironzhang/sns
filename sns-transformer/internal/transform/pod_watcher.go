package transform

import (
	"context"

	"k8s.io/client-go/tools/cache"

	"github.com/ironzhang/tlog"

	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/pkg/snsutil"
	"github.com/ironzhang/sns/sns-transformer/internal/update"
)

type podWatcher struct {
	targetNamespace string
	cmnb            *snsutil.ClusterMetaNameBuilder
	updater         *update.Updater
}

func (p *podWatcher) OnWatch(indexer cache.Indexer, event k8sclient.Event) error {
	tlog.Debugw("on watch", "action", event.Action, "key", event.Key)

	objects, err := indexer.Index("cluster_index", event.Object)
	if err != nil {
		tlog.Warnw("list objects by cluster index", "object", event.Object, "error", err)
		return nil
	}

	clusters := objectsToClusters(p.cmnb, p.targetNamespace, objects)
	if len(clusters) <= 0 {
		cnames, err := objectToCNames(p.cmnb, event.Object)
		if err != nil {
			return nil
		}
		if len(cnames) <= 0 {
			return nil
		}

		tlog.Debugw("delete clusters", "namespace", p.targetNamespace, "cnames", cnames)
		if err = p.updater.DeleteClusters(context.Background(), p.targetNamespace, cnames); err != nil {
			tlog.Errorw("delete clusters", "cnames", cnames, "error", err)
			return err
		}
		return nil
	}

	tlog.Debugw("update clusters", "clusters", clusters)
	if err = p.updater.UpdateClusters(context.Background(), clusters); err != nil {
		tlog.Errorw("update clusters", "clusters", clusters, "error", err)
		return err
	}
	return nil
}
