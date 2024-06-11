package agent

import (
	"context"
	"time"

	"go.uber.org/multierr"

	"github.com/ironzhang/sns/sns-agent/internal/watch"
)

type Agent struct {
	watcher *watch.Watcher
}

// New returns an instance of Agent.
func New(w *watch.Watcher) *Agent {
	return &Agent{
		watcher: w,
	}
}

func (p *Agent) WatchDomains(ctx context.Context, domains []string, ttl time.Duration) error {
	var reterr error
	for _, domain := range domains {
		err := p.watcher.WatchDomain(ctx, domain, ttl)
		if err != nil {
			reterr = multierr.Append(reterr, err)
		}
	}
	return reterr
}

func (p *Agent) ListWatchDomains(ctx context.Context) []string {
	return p.watcher.ListWatchDomains(ctx)
}
