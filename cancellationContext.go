package goCommsDefinitions

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
)

type ICancellationContext interface {
	Add(f func()) error
	Cancel()
	CancelWithError(err error)
}

type cancellationContext struct {
	cancelFunc    context.CancelFunc
	cancelContext context.Context
	logger        *zap.Logger
	f             []func()
	cancelCalled  bool
	closer        io.Closer
}

func (self *cancellationContext) CancelWithError(err error) {
	self.Cancel()
}

func (self *cancellationContext) Add(f func()) error {
	self.f = append(self.f, f)
	return nil
}

func (self *cancellationContext) CancelContext() context.Context {
	return self.cancelContext
}

func (self *cancellationContext) CancelFunc() context.CancelFunc {
	return self.Cancel
}

func (self *cancellationContext) Cancel() {
	if !self.cancelCalled {
		self.cancelCalled = true
		self.logger.Info(fmt.Sprintf("Cancel func for connection called"))
		self.cancelFunc()
		if self.closer != nil {
			self.closer.Close()
		}
		for _, f := range self.f {
			if f != nil {
				f()
			}
		}
	}
}

func NewCancellationContext(
	cancelFunc context.CancelFunc,
	cancelContext context.Context,
	logger *zap.Logger,
	closer io.Closer,
) *cancellationContext {
	return &cancellationContext{
		cancelFunc:    cancelFunc,
		cancelContext: cancelContext,
		logger:        logger,
		cancelCalled:  false,
		closer:        closer,
	}
}
