package RoTracing

import "github.com/opentracing/opentracing-go"

type RoogooTracer struct {
}

func (t RoogooTracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return nil
}
func (t RoogooTracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return nil
}
func (t RoogooTracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {

	return nil, nil
}
