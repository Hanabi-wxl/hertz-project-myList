// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz-mylist/biz/handler/myHandler"
	"hertz-mylist/biz/router/middleware"
	"hertz-mylist/conf"
)

func main() {
	conf.InitService()

	h := server.Default()

	middleware.JwtMwInit()

	h.Use(middleware.CorsMw())

	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(myHandler.Handler)))

	register(h)
	h.Spin()
}
