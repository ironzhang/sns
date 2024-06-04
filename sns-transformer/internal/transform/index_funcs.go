package transform

import (
	"errors"

	corev1 "k8s.io/api/core/v1"

	"github.com/ironzhang/tlog"
)

func appIndexFunc(obj interface{}) ([]string, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		tlog.Errorw("object is not a pod", "object", obj)
		return nil, errors.New("object is not a pod")
	}
	return []string{pod.ObjectMeta.Labels["app"]}, nil
}
