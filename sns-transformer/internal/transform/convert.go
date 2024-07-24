package transform

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ironzhang/tlog"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
	"github.com/ironzhang/sns/pkg/k8sutil"
	"github.com/ironzhang/sns/pkg/snsutil"
)

type podCollection struct {
	namespace string
	clusters  map[string]*coresnsv1.SNSCluster
}

func (p *podCollection) getOrNewCluster(n snsutil.ClusterMetadataName) *coresnsv1.SNSCluster {
	cname := n.String()
	c, ok := p.clusters[cname]
	if !ok {
		c = &coresnsv1.SNSCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      cname,
				Namespace: p.namespace,
				Labels: map[string]string{
					"cluster": n.ClusterName,
					"domain":  n.Domain(),
				},
			},
			Spec: coresnsv1.ClusterSpec{
				Kind:      snsutil.K8S_ClusterKind,
				Endpoints: make([]coresnsv1.Endpoint, 0),
			},
		}
		p.clusters[cname] = c
	}
	return c
}

func (p *podCollection) AddPod(pod *corev1.Pod) {
	if pod.Status.PodIP == "" {
		return
	}
	if pod.ObjectMeta.Labels["app"] == "" {
		return
	}

	for _, c := range pod.Spec.Containers {
		for _, port := range c.Ports {
			if port.Name == "" {
				continue
			}

			cmn := snsutil.NewClusterMetadataName(pod.ObjectMeta.Labels["cluster"], port.Name, pod.ObjectMeta.Labels["app"])
			cluster := p.getOrNewCluster(cmn)
			cluster.Spec.Endpoints = append(cluster.Spec.Endpoints, coresnsv1.Endpoint{
				Addr:   snsutil.JoinHostPort(pod.Status.PodIP, int(port.ContainerPort)),
				State:  k8sutil.GetPodState(pod),
				Weight: k8sutil.GetPodWeight(pod),
				Tags:   k8sutil.GetPodTags(pod),
			})
		}
	}
}

func (p *podCollection) ListClusters() []coresnsv1.SNSCluster {
	clusters := make([]coresnsv1.SNSCluster, 0, len(p.clusters))
	for _, cluster := range p.clusters {
		clusters = append(clusters, *cluster)
	}
	return clusters
}

func objectsToClusters(namespace string, objects []interface{}) []coresnsv1.SNSCluster {
	pc := podCollection{
		namespace: namespace,
		clusters:  make(map[string]*coresnsv1.SNSCluster),
	}
	for _, object := range objects {
		pod, ok := object.(*corev1.Pod)
		if !ok {
			tlog.Errorw("object is not a pod", "object", object)
			continue
		}
		pc.AddPod(pod)
	}
	return pc.ListClusters()
}

func objectToCNames(object interface{}) ([]string, error) {
	pod, ok := object.(*corev1.Pod)
	if !ok {
		tlog.Errorw("object is not a pod", "object", object)
		return nil, errors.New("object is not a pod")
	}
	if pod.ObjectMeta.Labels["app"] == "" {
		return nil, nil
	}

	cnames := make([]string, 0)
	for _, c := range pod.Spec.Containers {
		for _, port := range c.Ports {
			if port.Name == "" {
				continue
			}
			cmn := snsutil.NewClusterMetadataName(pod.ObjectMeta.Labels["cluster"], port.Name, pod.ObjectMeta.Labels["app"])
			cnames = append(cnames, cmn.String())
		}
	}
	return cnames, nil
}
