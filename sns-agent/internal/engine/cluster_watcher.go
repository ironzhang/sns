package engine

import (
	"encoding/json"

	"k8s.io/client-go/tools/cache"

	"github.com/ironzhang/superlib/superutil/supermodel"
	"github.com/ironzhang/tlog"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/filewrite"
	"github.com/ironzhang/sns/pkg/k8sclient"
	"github.com/ironzhang/sns/pkg/superconv"
	"github.com/ironzhang/sns/sns-agent/internal/paths"
)

type clusterWatcher struct {
	sorter  *clusterSorter
	pathmgr *paths.PathManager
	fwriter *filewrite.FileWriter
}

func (p *clusterWatcher) OnWatch(indexer cache.Indexer, event k8sclient.Event) error {
	tlog.Debugw("on watch", "event", event)

	c, ok := event.Object.(*coresnsv1.SNSCluster)
	if !ok {
		tlog.Errorw("object is not a cluster", "object", event.Object)
		return nil
	}
	return p.refresh(indexer, c)
}

func (p *clusterWatcher) OnRefresh(indexer cache.Indexer) {
	for _, obj := range indexer.List() {
		c, ok := obj.(*coresnsv1.SNSCluster)
		if !ok {
			tlog.Errorw("object is not a cluster", "object", obj)
			continue
		}

		tlog.Debugw("on refresh", "domain", c.ObjectMeta.Labels["domain"])
		p.refresh(indexer, c)
		return
	}
}

func (p *clusterWatcher) refresh(indexer cache.Indexer, c *coresnsv1.SNSCluster) error {
	model := supermodel.ServiceModel{
		Domain:   c.ObjectMeta.Labels["domain"],
		Clusters: p.objectsToClusters(indexer.List()),
	}
	err := p.writeModel(model)
	if err != nil {
		tlog.Errorw("write service model", "model", model, "error", err)
		return err
	}
	return nil
}

func (p *clusterWatcher) writeModel(m supermodel.ServiceModel) error {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	path := p.pathmgr.ServiceModelPath(m.Domain)
	if err = p.fwriter.WriteFile(path, data); err != nil {
		return err
	}
	return nil
}

func (p *clusterWatcher) objectsToClusters(objects []interface{}) []supermodel.Cluster {
	clusters := make([]supermodel.Cluster, 0, len(objects))
	for _, obj := range objects {
		c, ok := obj.(*coresnsv1.SNSCluster)
		if !ok {
			tlog.Errorw("object is not a cluster", "obj", obj)
			continue
		}
		clusters = append(clusters, superconv.ToSupermodelCluster(*c))
	}

	p.sorter.Sort(clusters)
	return clusters
}
