package ginSpan

//
//func RpcSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (spanctx context.Context, span trace.Span) {
//
//	md, b := metadata.FromIncomingContext(ctx)
//	if !b {
//		return ctx, nil
//	}
//	//value := ctx.Value(TracerKey)
//	//tracer, ok := value.(trace.Tracer)
//	//if !ok {
//	//	fmt.Println(" value.(RoTracing.Tracer) !ok")
//	//	return ctx, nil
//	//}
//
//	// gin 特殊
//	if c, ok := ctx.(*gin.Context); ok {
//		spanctx, span = tracer.Start(c.Request.Context(), spanName, opts...)
//
//		/*
//			在这里每次注入新的 Attr
//			1. host
//		*/
//		// 1. 从 context 中获取 "public attr"
//		// attr:=ctx.Value("")
//		// 2. 注入 public attr
//		// span.SetAttributes(attr)
//
//		spanctx = context.WithValue(spanctx, TracerKey, tracer)
//		// return spanctx, span
//	} else {
//		spanctx, span = tracer.Start(ctx, spanName, opts...)
//	}
//
//	// 设置 Attr
//	attrkv, ok := ctx.Value("attrkv").(map[string]string)
//	if ok {
//		SpanSetStringAttr(span, attrkv)
//	}
//
//	SpanSetStringAttr(span, map[string]string{
//		"server.host": os.Getenv("HOSTNAME"),
//	})
//
//	return spanctx, span
//}
