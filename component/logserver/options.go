package logserver

import (
	"context"
	"eventgo/component/logserver/modules/logconsumer"
	"eventgo/core/store/redis"
)

type Options struct {
	Rds         *redis.Store
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	Context context.Context
	cancel  context.CancelFunc

	logConsumer logconsumer.LogConsumer

	Signal bool
}

type JournalHandle func(msg []byte)

func newOptions(opts ...Option) Options {
	opt := Options{
		Context: context.Background(),
		Signal:  true,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func Context1(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func HandleSignal(b bool) Option {
	return func(o *Options) {
		o.Signal = b
	}
}

func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
