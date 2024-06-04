package main

import (
	"os"

	"github.com/ironzhang/superlib/ctxutil"
	"github.com/ironzhang/tapp"
	"github.com/ironzhang/tlog/zaplog"

	"github.com/ironzhang/sns/sns-transformer/internal/app"
)

var (
	Version   = "Unknown"
	GitCommit = "Unknown"
	BuildTime = "Unknown"
)

func main() {
	a := &app.Application{}
	f := tapp.Framework{
		Version: &tapp.VersionInfo{
			Version:   Version,
			GitCommit: GitCommit,
			BuildTime: BuildTime,
		},
		Application:       a,
		Config:            app.Conf,
		Runners:           []tapp.RunFunc{a.Run},
		LoggerContextHook: zaplog.ContextHookFunc(ctxutil.ContextHook),
	}
	f.Main(os.Args)
}
