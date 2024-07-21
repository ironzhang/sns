package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/ironzhang/superlib/httputils/echoutil"
	"github.com/ironzhang/superlib/httputils/echoutil/echorpc"
	"github.com/ironzhang/superlib/httputils/httpclient"
	"github.com/ironzhang/superlib/httputils/httpclient/interceptors"
	"github.com/ironzhang/supernamego"
	"github.com/labstack/echo"
)

func SupernameResolve(ctx context.Context, addr string) (string, error) {
	host, _, err := supernamego.Lookup(ctx, addr)
	return host, err
}

var client = httpclient.Client{
	Addr: "sns/http.callee.example",
	Client: http.Client{
		Timeout: time.Second,
	},
	Resolve: SupernameResolve,
	Interceptors: []httpclient.Interceptor{
		interceptors.AccessLogInterceptor(),
	},
}

func HandleCall(ctx context.Context, in interface{}, out *string) error {
	return client.Post(context.TODO(), "/echo", nil, "hello, world", out)
}

func main() {
	var addr string
	flag.StringVar(&addr, "callee-addr", "sns/http.callee", "the callee address")
	flag.Parse()
	client.Addr = addr

	e := echo.New()
	e.HTTPErrorHandler = echoutil.HTTPErrorHandler
	e.Use(echoutil.AccessLogMiddleware())
	e.GET("/call", echorpc.HandlerFunc(HandleCall))
	e.Start(":8000")
}
