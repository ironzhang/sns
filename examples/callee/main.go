package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ironzhang/superlib/httputils/echoutil"
	"github.com/ironzhang/superlib/httputils/echoutil/echorpc"
	"github.com/labstack/echo"
)

func HandleEcho(ctx context.Context, in string, out *string) error {
	name, _ := os.Hostname()
	*out = fmt.Sprintf("%s: %s", name, in)
	return nil
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = echoutil.HTTPErrorHandler
	e.Use(echoutil.AccessLogMiddleware())
	e.POST("/echo", echorpc.HandlerFunc(HandleEcho))
	e.Start(":8001")
}
