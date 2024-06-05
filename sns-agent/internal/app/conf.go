package app

import "github.com/ironzhang/superlib/superutil/parameter"

type ListenConf struct {
	Addr string
}

type Config struct {
	Namespace    string
	ResourcePath string
	Listen       ListenConf
}

var Conf = &Config{
	Namespace:    "sns",
	ResourcePath: parameter.Param.Watch.ResourcePath,
	Listen: ListenConf{
		Addr: parameter.Param.Agent.Server,
	},
}