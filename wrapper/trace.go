package wrapper

import (
	"context"
	"fmt"
	"github.com/Gitforxuyang/microBase/trace"
	"github.com/Gitforxuyang/microBase/util"
	"github.com/opentracing/opentracing-go"
)

type RedisSpan interface {
	Finish(err error)
	SetCmd(cmd string)
}

type redisSpan struct {
	span opentracing.Span
}

func (r *redisSpan) Finish(err error) {
	if err != nil {
		r.span.SetTag("error", true)
		r.span.LogKV("error.object", err.Error(), "event", "error")
	}
	fmt.Printf("finish")
	r.span.Finish()
}
func (r *redisSpan) SetCmd(cmd string) {
	r.span.SetTag("db.statement", cmd)
}

type RedisTraceHandleFunc func(ctx context.Context, cmd string) RedisSpan

func RedisTraceWrapper(ot opentracing.Tracer, instance string) RedisTraceHandleFunc {
	return func(ctx context.Context, cmd string) RedisSpan {
		ctx, span, err := trace.StartSpanFromContext(ctx, ot, cmd)
		span.SetTag("span.kind", "client")
		span.SetTag("component", "microRedis")
		span.SetTag("db.instance", instance)
		span.SetTag("db.type", "redis")
		rspan := &redisSpan{span: span}
		if err != nil {
			util.Error(ctx, err.Error())
			rspan.span = nil
		}
		return rspan
	}
}
