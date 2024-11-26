package Rootel

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var globalTraceProvider trace.TracerProvider

func Register(appName string, endpoint string) (gin.HandlerFunc, trace.TracerProvider) {

	opts := []otelgin.Option{
		ProviderOption(appName, endpoint),
		PropagationExtractOption(),
	}

	return otelgin.Middleware(appName, opts...), globalTraceProvider
}

func ProviderOption(appName string, endpoint string) otelgin.Option {
	// 1. 注册 Provider
	provider, err := initProvider(appName, endpoint)
	if err != nil {
		panic(err)
	}
	// 全局注册Provider
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	SetPd(&provider)
	return otelgin.WithTracerProvider(provider)
}

func RegisterGrpcProvider(appName string, endpoint string) trace.TracerProvider {
	//if otel.GetTracerProvider() != nil {
	//	fmt.Println(otel.GetTracerProvider())
	//	return otel.GetTracerProvider()
	//}
	return GrpcProviderOption(appName, endpoint)
}

func GrpcProviderOption(appName string, endpoint string) trace.TracerProvider {

	// 1. 注册 Provider
	provider, err := initProvider(appName, endpoint)
	if err != nil {
		panic(err)
	}
	// 全局注册Provider
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	//return otelgrpc.WithTracerProvider(provider)
	return otel.GetTracerProvider()
}

func SetPd(pd *trace.TracerProvider) {
	globalTraceProvider = *pd

}
