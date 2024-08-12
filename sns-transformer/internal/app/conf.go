package app

import (
	"k8s.io/client-go/tools/clientcmd"

	"github.com/ironzhang/sns/sns-transformer/internal/transform"
)

type KubeConf struct {
	KubeConfigPath string
}

type Config struct {
	Kube      KubeConf
	Transform transform.Options
}

var Conf = &Config{
	Kube: KubeConf{
		KubeConfigPath: clientcmd.RecommendedHomeFile,
	},
	Transform: transform.Options{
		SourceNamespace: "default",
		TargetNamespace: "sns",
		DefaultZone:     "az00",
		DefaultLane:     "default",
	},
}
