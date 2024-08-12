package app

import (
	"github.com/ironzhang/sns/sns-agent/internal/engine"
	"github.com/ironzhang/superlib/superutil/parameter"
	"k8s.io/client-go/tools/clientcmd"
)

type ListenConf struct {
	Addr string
}

type KubeConf struct {
	KubeConfigPath string
}

type Config struct {
	Listen       ListenConf
	Kube         KubeConf
	Engine       engine.Options
	ResourcePath string
}

var Conf = &Config{
	Listen: ListenConf{
		Addr: parameter.Param.Agent.Server,
	},
	Kube: KubeConf{
		KubeConfigPath: clientcmd.RecommendedHomeFile,
	},
	Engine: engine.Options{
		Namespace:   "sns",
		DefaultZone: "az00",
		DefaultLane: "default",
		DefaultKind: "k8s",
	},
	ResourcePath: parameter.Param.Watch.ResourcePath,
}
