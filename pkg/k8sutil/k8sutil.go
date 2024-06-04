package k8sutil

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetObjectMeta(object interface{}) *metav1.ObjectMeta {
	m, ok := object.(*metav1.ObjectMeta)
	if !ok {
		return nil
	}
	return m
}
