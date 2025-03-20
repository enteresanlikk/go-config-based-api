package http

import (
	"github.com/goccy/go-json"

	"github.com/valyala/fasthttp"
)

func SetJsonResponseHeader(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
}

func SetResponseStatus(ctx *fasthttp.RequestCtx, status int) {
	ctx.SetStatusCode(status)
}

func SetJsonResponseBody(ctx *fasthttp.RequestCtx, data interface{}) error {
	return json.NewEncoder(ctx).Encode(data)
}

func JsonResponse(ctx *fasthttp.RequestCtx, data interface{}) error {
	SetJsonResponseHeader(ctx)
	return SetJsonResponseBody(ctx, data)
}

func JsonResponseWithStatus(ctx *fasthttp.RequestCtx, status int, data interface{}) error {
	SetJsonResponseHeader(ctx)
	SetResponseStatus(ctx, status)
	return SetJsonResponseBody(ctx, data)
}
