package transform

import (
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/ironzhang/tlog"
)

func clusterIndexFunc(obj interface{}) ([]string, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		tlog.Errorw("object is not a pod", "object", obj)
		return nil, errors.New("object is not a pod")
	}

	clusterID := fmt.Sprintf("%s.%s", pod.ObjectMeta.Labels["cluster"], pod.ObjectMeta.Labels["app"])
	return []string{clusterID}, nil
}
