package service

import (
	"context"

	api "github.com/Ccmuyu/biz-demo/gomall/thrift_api/hertz_gen/api"
	"github.com/cloudwego/hertz/pkg/app"
)

type GreetService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGreetService(Context context.Context, RequestContext *app.RequestContext) *GreetService {
	return &GreetService{RequestContext: RequestContext, Context: Context}
}

func (h *GreetService) Run(req *api.GreetReq) (resp *api.GreetResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
