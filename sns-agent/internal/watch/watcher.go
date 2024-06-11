package watch

import (
	"context"
	"time"

	"github.com/ironzhang/tlog"

	"github.com/ironzhang/sns/sns-agent/internal/engine"
	"github.com/ironzhang/sns/sns-agent/internal/ready"
)

// A Watcher watches domains.
type Watcher struct {
	notary     *notary
	engine     *engine.Engine
	inspection *ready.Inspection
}

// NewWatcher returns an instance of Watcher.
func NewWatcher(eng *engine.Engine, ins *ready.Inspection) *Watcher {
	return &Watcher{
		notary:     newNotary(),
		engine:     eng,
		inspection: ins,
	}
}

// WatchDomain watch domain in ttl duration.
//
// if ttl <= 0, means forever.
func (p *Watcher) WatchDomain(ctx context.Context, domain string, ttl time.Duration) error {
	do := func(ctx context.Context) (context.CancelFunc, error) {
		logger := tlog.WithContext(ctx).WithArgs("domain", domain)
		ctx, cancel := context.WithCancel(ctx)

		logger.Infow("watch domain")
		err := p.engine.WatchDomain(ctx, domain)
		if err != nil {
			logger.Errorw("watch domain", "error", err)
			return func() {}, err
		}

		return func() {
			logger.Infow("cancel watch domain")
			cancel()
		}, nil
	}

	_, err := p.notary.NewLease(ctx, domain, ttl, do)
	if err != nil {
		return err
	}

	if !p.inspection.ServiceReady(domain) {
		p.engine.RefreshClusters(ctx, domain)
	}
	//	if !p.inspection.RouteReady(domain) {
	//		p.engine.RefreshRoutes(ctx, domain)
	//	}

	return nil
}

// ListWatchDomains returns the domains which are watching.
func (p *Watcher) ListWatchDomains(ctx context.Context) []string {
	return p.notary.ListLeaseKeys(ctx)
}
