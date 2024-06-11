package protocol

import "github.com/ironzhang/superlib/timeutil"

// WatchDomainsReq is a request for watching domains.
type WatchDomainsReq struct {
	Domains      []string          // the domain list that require to watch
	TTL          timeutil.Duration // time to live, <= 0 means forever
	Asynchronous bool              // asynchronous call
}

// ListWatchDomainsResp is a response for listing watch domains.
type ListWatchDomainsResp struct {
	Domains []string // the domain list that the agent is subscribing
}
