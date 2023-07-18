package internal

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dreamyourlife/lscdl/pkg/kafka"
	"github.com/dreamyourlife/lscdl/pkg/pg"
	"github.com/dreamyourlife/lscdl/pkg/rabbitmq"
)

var (
	ctx           context.Context
	ctxCancelFunc context.CancelFunc
)

func initContext() {
	ctx, ctxCancelFunc = context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		for range c {
			arrowCtxCancelFunc()
		}
	}()
	//kafka.InitContext(arrowCtx)
}
