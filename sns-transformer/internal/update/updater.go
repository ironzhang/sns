package update

import (
	"context"

	"go.uber.org/multierr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ironzhang/tlog"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
	coresnsv1client "github.com/ironzhang/sns/kernel/clients/coresnsclients/clientset/versioned/typed/core.sns.io/v1"
)

type Updater struct {
	clustersGetter coresnsv1client.SNSClustersGetter
}

func NewUpdater(g coresnsv1client.SNSClustersGetter) *Updater {
	return &Updater{
		clustersGetter: g,
	}
}

func (p *Updater) UpdateCluster(ctx context.Context, c coresnsv1.SNSCluster) error {
	iface := p.clustersGetter.SNSClusters(c.ObjectMeta.Namespace)

	exist, err := iface.Get(ctx, c.ObjectMeta.Name, metav1.GetOptions{})
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return err
		}
		if _, err = iface.Create(ctx, &c, metav1.CreateOptions{}); err != nil {
			tlog.WithContext(ctx).Errorw("create", "cluster", c, "error", err)
			return err
		}
		return nil
	}

	c.ObjectMeta.ResourceVersion = exist.ObjectMeta.ResourceVersion
	if c.Spec.Tags == nil {
		c.Spec.Tags = exist.Spec.Tags
	}
	if c.Spec.Endpoints == nil {
		c.Spec.Endpoints = exist.Spec.Endpoints
	}

	if _, err = iface.Update(ctx, &c, metav1.UpdateOptions{}); err != nil {
		tlog.WithContext(ctx).Errorw("update", "cluster", c, "error", err)
		return err
	}
	return nil
}

func (p *Updater) DeleteCluster(ctx context.Context, namespace, cname string) error {
	iface := p.clustersGetter.SNSClusters(namespace)

	err := iface.Delete(ctx, cname, metav1.DeleteOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return nil
		}
		tlog.WithContext(ctx).Errorw("delete", "cluster", cname, "error", err)
		return err
	}
	return nil
}

func (p *Updater) UpdateClusters(ctx context.Context, clusters []coresnsv1.SNSCluster) (err error) {
	for _, c := range clusters {
		err = multierr.Append(err, p.UpdateCluster(ctx, c))
	}
	return err
}

func (p *Updater) DeleteClusters(ctx context.Context, namespace string, cnames []string) (err error) {
	for _, cname := range cnames {
		err = multierr.Append(err, p.DeleteCluster(ctx, namespace, cname))
	}
	return err
}
