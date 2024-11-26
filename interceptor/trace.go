package interceptor

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	//TraceIDKey = "Trace-ID"
	TraceIDKey = "Rootel-go-contrib-tracer"
)

func TracingSeverInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("==============server interceptor=======================")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("metadata is required")
		}
		fmt.Println("=====md:========", md)
		traceID := md.Get(TraceIDKey)
		if len(traceID) == 0 {
			return nil, fmt.Errorf("RoTracing-id is required")
		}

		return handler(ctx, req)
	}
}

// TracingClientInterceptor 返回一个客户端拦截器，用于链路追踪
func TracingClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		value := ctx.Value(TraceIDKey)
		_, ok := value.(trace.Tracer)
		if !ok {
			fmt.Println(" value.(RoTracing.Tracer) !ok")
			return errors.New("value.(RoTracing.Tracer) !ok")
		}
		traceID := "AAAAAAA11199999999"
		// 将 RoTracing ID 添加到请求的元数据中
		md := metadata.Pairs(TraceIDKey, traceID)
		ctx = metadata.NewOutgoingContext(ctx, md)

		fmt.Printf("Sending request with RoTracing ID: %s\n", traceID)

		fmt.Println("ctx:", ctx)
		//// 调用实际的方法
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return err
		}

		//fmt.Printf("Received response with RoTracing ID: %s\n", traceID)
		return nil
	}
}
