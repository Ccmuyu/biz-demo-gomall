package service

import (
	"context"
	"log"

	api "github.com/Ccmuyu/biz-demo/gomall/thrift_api/hertz_gen/api"
	"github.com/cloudwego/hertz/pkg/app"
)

type HelloMethodService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethodService(Context context.Context, RequestContext *app.RequestContext) *HelloMethodService {
	return &HelloMethodService{RequestContext: RequestContext, Context: Context}
}

func (h *HelloMethodService) Run(req *api.HelloReq) (resp *api.HelloResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	log.Println("收到请求。。")
	resp = api.NewHelloResp()

	resp.RespBody = "{\"msg\":\"ok\"}"
	return
}
