package ginSpan

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	/*
		info:
		https://github.com/open-telemetry/opentelemetry-go-contrib/blob/849072ef827b4abab754253e1e63e7b410a31084/instrumentation/github.com/gin-gonic/gin/otelgin/gintrace.go#L42
		https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/gintrace.go

	*/
	TracerKey = "otel-go-contrib-tracer"
)

func Span(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (spanctx context.Context, span trace.Span) {

	value := ctx.Value(TracerKey)
	tracer, ok := value.(trace.Tracer)
	if !ok {
		fmt.Println(" value.(RoTracing.Tracer) !ok")
		return ctx, nil
	}

	// gin 特殊
	if c, ok := ctx.(*gin.Context); ok {
		spanctx, span = tracer.Start(c.Request.Context(), spanName, opts...)

		/*
			在这里每次注入新的 Attr
			1. host
		*/
		// 1. 从 context 中获取 "public attr"
		// attr:=ctx.Value("")
		// 2. 注入 public attr
		// span.SetAttributes(attr)

		spanctx = context.WithValue(spanctx, TracerKey, tracer)
		// return spanctx, span
	} else {
		spanctx, span = tracer.Start(ctx, spanName, opts...)
	}

	// 设置 Attr
	attrkv, ok := ctx.Value("attrkv").(map[string]string)
	if ok {
		SpanSetStringAttr(span, attrkv)
	}

	SpanSetStringAttr(span, map[string]string{
		"server.host": os.Getenv("HOSTNAME"),
	})

	return spanctx, span
}

func SpanSetStringAttr(span trace.Span, kvs map[string]string) {
	attrkv := []attribute.KeyValue{}

	for k, v := range kvs {
		attrkv = append(attrkv, attribute.KeyValue{
			Key:   attribute.Key(k),
			Value: attribute.StringValue(v),
		})
	}

	span.SetAttributes(attrkv...)
}

func SpanContextWithAttr(ctx context.Context, kv map[string]string) context.Context {

	value := ctx.Value("attrkv")
	attrkv, ok := value.(map[string]string)
	if !ok {
		attrkv = make(map[string]string, 0)
	}

	for k, v := range kv {
		attrkv[k] = v
	}

	return context.WithValue(ctx, "attrkv", attrkv)
}
