package goCommsDefinitions

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideCancelContext(cancelContext context.Context) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Target: func(
				params struct {
					fx.In
					Logger         *zap.Logger
					ConnectionName string `name:"ConnectionName"`
				},
			) (context.Context, context.CancelFunc, ICancellationContext, error) {
				ctx, cancelFunc := context.WithCancel(cancelContext)
				cancelInstance := NewCancellationContext(params.ConnectionName, cancelFunc, ctx, params.Logger, nil)
				return ctx, cancelInstance.Cancel, cancelInstance, nil
			},
		},
	)
}
