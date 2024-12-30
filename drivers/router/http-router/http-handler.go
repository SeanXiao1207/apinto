package http_router

import (
	"net/http"
	"time"

	"github.com/eolinker/apinto/entries/ctx_key"

	http_service "github.com/eolinker/apinto/node/http-context"

	http_complete "github.com/eolinker/apinto/drivers/router/http-router/http-complete"
	"github.com/eolinker/apinto/service"

	"github.com/eolinker/eosc/eocontext"
	http_context "github.com/eolinker/eosc/eocontext/http-context"
)

var completeCaller = http_complete.NewHttpCompleteCaller()

type httpHandler struct {
	completeHandler eocontext.CompleteHandler

	routerName  string
	routerId    string
	serviceName string
	finisher    eocontext.FinishHandler
	service     service.IService
	filters     eocontext.IChainPro
	disable     bool
	websocket   bool
	labels      map[string]string
	retry       int
	timeout     time.Duration
}

func (h *httpHandler) Serve(ctx eocontext.EoContext) {
	httpContext, err := http_context.Assert(ctx)
	if err != nil {
		return
	}
	if h.disable {
		httpContext.Response().SetStatus(http.StatusNotFound, "")
		httpContext.Response().SetBody([]byte("router disable"))
		httpContext.FastFinish()
		return
	}
	if h.websocket {
		wsCtx, err := http_service.NewWebsocketContext(httpContext)
		if err != nil {
			httpContext.Response().SetStatus(http.StatusInternalServerError, "")
			httpContext.Response().SetBody([]byte(err.Error()))
			httpContext.FastFinish()
			return
		}
		ctx = wsCtx
	}

	for key, value := range h.labels {
		// 设置标签
		ctx.SetLabel(key, value)
	}
	//set retry timeout
	ctx.WithValue(ctx_key.CtxKeyRetry, h.retry)
	ctx.WithValue(ctx_key.CtxKeyTimeout, h.timeout)

	//Set Label
	ctx.SetLabel("api", h.routerName)
	ctx.SetLabel("api_id", h.routerId)
	ctx.SetLabel("service", h.serviceName)
	if h.service != nil {
		ctx.SetLabel("service_id", h.service.Id())
		ctx.SetLabel("service_title", h.service.Title())
	}

	ctx.SetLabel("method", httpContext.Request().Method())
	ctx.SetLabel("path", httpContext.Request().URI().RequestURI())
	ctx.SetLabel("ip", httpContext.Request().RealIp())

	ctx.SetCompleteHandler(h.completeHandler)
	ctx.SetBalance(h.service)
	ctx.SetUpstreamHostHandler(h.service)
	ctx.SetFinish(h.finisher)
	h.filters.Chain(ctx, completeCaller)
}
