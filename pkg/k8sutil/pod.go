package k8sutil

import (
	"strconv"

	corev1 "k8s.io/api/core/v1"

	coresnsv1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
)

func GetPodState(pod *corev1.Pod) coresnsv1.State {
	if pod.Status.Phase != corev1.PodRunning {
		return coresnsv1.Disabled
	}
	return coresnsv1.Enabled
}

func GetPodWeight(pod *corev1.Pod) int {
	s, ok := pod.ObjectMeta.Annotations["weight"]
	if !ok {
		return 100
	}
	weight, err := strconv.Atoi(s)
	if err != nil {
		return 100
	}
	if weight < 0 {
		return 0
	}
	return weight
}

func GetPodTags(pod *corev1.Pod) map[string]string {
	return nil
}
