package superconv

import (
	"github.com/ironzhang/superlib/superutil/supermodel"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
)

// ToSupermodelEndpoint convert coresnsv1.Endpoint to supermodel.Endpoint.
func ToSupermodelEndpoint(ep coresnsv1.Endpoint) supermodel.Endpoint {
	return supermodel.Endpoint{
		Addr:   ep.Addr,
		State:  supermodel.State(ep.State),
		Weight: ep.Weight,
	}
}

// ToSupermodelCluster convert coresnsv1.SNSCluster to supermodel.Cluster.
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

// ToSupermodelToken convert coresnsv1.Token to supermodel.Token.
func ToSupermodelToken(t coresnsv1.Token) supermodel.Token {
	return supermodel.Token{
		Type:   supermodel.TokenType(t.Type),
		Table:  t.Table,
		Key:    t.Key,
		Consts: t.Consts,
	}
}

// ToSupermodelRequirement convert coresnsv1.Requirement to supermodel.Requirement.
func ToSupermodelRequirement(r coresnsv1.Requirement) supermodel.Requirement {
	return supermodel.Requirement{
		Not:      r.Not,
		Operator: supermodel.Operator(r.Operator),
		Left:     ToSupermodelToken(r.Left),
		Right:    ToSupermodelToken(r.Right),
	}
}

// ToSupermodelLabelSelector convert coresnsv1.LabelSelector to supermodel.LabelSelector.
func ToSupermodelLabelSelector(s coresnsv1.LabelSelector) supermodel.LabelSelector {
	results := make(supermodel.LabelSelector, 0, len(s))
	for _, r := range s {
		results = append(results, ToSupermodelRequirement(r))
	}
	return results
}

// ToSupermodelRoutePolicy convert coresnsv1.SNSRoutePolicy to supermodel.RoutePolicy.
func ToSupermodelRoutePolicy(p coresnsv1.SNSRoutePolicy) supermodel.RoutePolicy {
	selectors := make([]supermodel.LabelSelector, 0, len(p.Spec.LabelSelectors))
	for _, s := range p.Spec.LabelSelectors {
		selectors = append(selectors, ToSupermodelLabelSelector(s))
	}
	return supermodel.RoutePolicy{
		EnableScript:   p.Spec.RouteScript.Enable,
		LabelSelectors: selectors,
	}
}
