package jaeger

import (
	"io"

	"github.com/AssylzhanZharzhanov/connect/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

// NewJaegerTracer - creates a tracer
func NewJaegerTracer(cfg config.AppConfig) (opentracing.Tracer, io.Closer, error) {
	configurations := &jaegerCfg.Configuration{
		ServiceName: cfg.ServiceName,

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:           cfg.LogSpans,
			LocalAgentHostPort: cfg.HostPort,
		},
	}

	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

	return configurations.NewTracer(
		jaegerCfg.Logger(jaeger.StdLogger),
		jaegerCfg.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		jaegerCfg.Injector(opentracing.TextMap, zipkinPropagator),
		jaegerCfg.Injector(opentracing.Binary, zipkinPropagator),
		jaegerCfg.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
		jaegerCfg.Extractor(opentracing.TextMap, zipkinPropagator),
		jaegerCfg.Extractor(opentracing.Binary, zipkinPropagator),
		jaegerCfg.ZipkinSharedRPCSpan(false),
	)
}
