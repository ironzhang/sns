package superconv

import (
	"github.com/ironzhang/superlib/superutil/supermodel"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
)

// ToSupermodelEndpoint convert coresnsv1.Endpoint to supermodel.Endpoint
func ToSupermodelEndpoint(ep coresnsv1.Endpoint) supermodel.Endpoint {
	return supermodel.Endpoint{
		Addr:   ep.Addr,
		State:  supermodel.State(ep.State),
		Weight: ep.Weight,
	}
}

// ToSupermodelCluster convert coresnsv1.SNSCluster to supermodel.Cluster
func ToSupermodelCluster(c coresnsv1.SNSCluster) supermodel.Cluster {
	endpoints := make([]supermodel.Endpoint, 0, len(c.Spec.Endpoints))
	for _, ep := range c.Spec.Endpoints {
		endpoints = append(endpoints, ToSupermodelEndpoint(ep))
	}
	return supermodel.Cluster{
		Name:      c.ObjectMeta.Labels["cluster"],
		Labels:    c.Spec.Labels,
		Endpoints: endpoints,
	}
}

/*
// ToSupermodelDestination convert superdnsv1.Destination to supermodel.Destination
func ToSupermodelDestination(d superdnsv1.Destination) supermodel.Destination {
	return supermodel.Destination{
		Cluster: d.Cluster,
		Percent: d.Percent,
	}
}

// ToSupermodelRouteStrategy convert superdnsv1.Route to supermodel.RouteStrategy
func ToSupermodelRouteStrategy(r superdnsv1.Route) supermodel.RouteStrategy {
	dests := make([]supermodel.Destination, 0, len(r.Spec.DefaultDestinations))
	for _, d := range r.Spec.DefaultDestinations {
		dests = append(dests, ToSupermodelDestination(d))
	}
	return supermodel.RouteStrategy{
		EnableScript:        r.Spec.EnableScript,
		DefaultDestinations: dests,
	}
}
*/
