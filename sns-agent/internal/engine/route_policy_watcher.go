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

type routePolicyWatcher struct {
	pathmgr *paths.PathManager
	fwriter *filewrite.FileWriter
}

func (p *routePolicyWatcher) OnWatch(indexer cache.Indexer, event k8sclient.Event) error {
	tlog.Debugw("on watch", "event", event)

	rp, ok := event.Object.(*coresnsv1.SNSRoutePolicy)
	if !ok {
		tlog.Errorw("object is not a route policy", "object", event.Object)
		return nil
	}
	return p.refresh(rp)
}

func (p *routePolicyWatcher) OnRefresh(indexer cache.Indexer) {
	for _, obj := range indexer.List() {
		rp, ok := obj.(*coresnsv1.SNSRoutePolicy)
		if !ok {
			tlog.Errorw("object is not a route policy", "object", obj)
			continue
		}

		tlog.Debugw("on refresh", "domain", rp.ObjectMeta.Name)
		p.refresh(rp)
		return
	}
}

func (p *routePolicyWatcher) refresh(rp *coresnsv1.SNSRoutePolicy) error {
	model := supermodel.RouteModel{
		Domain: rp.ObjectMeta.Name,
		Policy: superconv.ToSupermodelRoutePolicy(*rp),
	}
	err := p.writeModel(model)
	if err != nil {
		tlog.Errorw("write route model", "model", model, "error", err)
		return err
	}
	err = p.writeScript(model.Domain, rp.Spec.RouteScript.Content)
	if err != nil {
		tlog.Errorw("write route script", "domain", model.Domain, "error", err)
		return err
	}
	return nil
}

func (p *routePolicyWatcher) writeModel(m supermodel.RouteModel) error {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	path := p.pathmgr.RouteModelPath(m.Domain)
	if err = p.fwriter.WriteFile(path, data); err != nil {
		return err
	}
	return nil
}

func (p *routePolicyWatcher) writeScript(domain, content string) error {
	path := p.pathmgr.RouteScriptPath(domain)
	if err := p.fwriter.WriteFile(path, []byte(content)); err != nil {
		return err
	}
	return nil
}
