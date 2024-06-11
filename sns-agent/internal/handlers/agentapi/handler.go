package agentapi

import (
	"context"
	"time"

	"github.com/ironzhang/superlib/httputils/echoutil/echorpc"
	"github.com/ironzhang/tlog"
	"github.com/labstack/echo"

	"github.com/ironzhang/sns/pkg/protocol"
	"github.com/ironzhang/sns/sns-agent/internal/agent"
)

// Handler is an echo rpc handler to handle agent api.
type Handler struct {
	agent *agent.Agent
}

// Register registers agent api.
func Register(e *echo.Echo, h *Handler) {
	e.POST("/sns/agent/api/v1/watch/domains", echorpc.HandlerFunc(h.WatchDomains))
	e.GET("/sns/agent/api/v1/list/watch/domains", echorpc.HandlerFunc(h.ListWatchDomains))
}

// NewHandler returns an instance of Handler.
func NewHandler(a *agent.Agent) *Handler {
	return &Handler{agent: a}
}

// WatchDomains handles the subscribe domains request.
func (p *Handler) WatchDomains(ctx context.Context, req *protocol.WatchDomainsReq, resp interface{}) error {
	fn := func() error {
		err := p.agent.WatchDomains(ctx, req.Domains, time.Duration(req.TTL))
		if err != nil {
			tlog.WithContext(ctx).Errorw("watch domains", "domains", req.Domains, "error", err)
			return err
		}
		return nil
	}

	if req.Asynchronous {
		go fn()
		return nil
	}
	return fn()
}

// ListWatchDomains handles the list subscribe domains request.
func (p *Handler) ListWatchDomains(ctx context.Context, req interface{}, resp *protocol.ListWatchDomainsResp) error {
	resp.Domains = p.agent.ListWatchDomains(ctx)
	return nil
}
