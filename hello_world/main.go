package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello,你好呀～"))
	})

	h.Spin()
}
