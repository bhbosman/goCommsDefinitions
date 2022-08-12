package goCommsDefinitions

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"sync"
)

type ICancellationContext interface {
	Add(connectionId string, f func()) error
	Remove(connectionId string) error
	Cancel()
	CancelWithError(err error)
}

type cancellationContext struct {
	mutex         sync.Mutex
	cancelFunc    context.CancelFunc
	cancelContext context.Context
	logger        *zap.Logger
	f             map[string]func()
	cancelCalled  bool
	closer        io.Closer
}

func (self *cancellationContext) Remove(connectionId string) error {
	if !self.cancelCalled {
		self.mutex.Lock()
		defer self.mutex.Unlock()
		delete(self.f, connectionId)
	}
	return nil
}

func (self *cancellationContext) CancelWithError(err error) {
	self.Cancel()
}

func (self *cancellationContext) Add(connectionId string, f func()) error {
	if !self.cancelCalled {
		self.mutex.Lock()
		defer self.mutex.Unlock()
		//
		if f, ok := self.f[connectionId]; ok {
			f()
		}
		self.f[connectionId] = f
	}
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
		self.mutex.Lock()
		fArray := make([]func(), 0, len(self.f))
		for _, f := range self.f {
			fArray = append(fArray, f)
		}
		self.f = make(map[string]func())
		self.mutex.Unlock()
		for _, f := range fArray {
			f()
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
		f:             make(map[string]func()),
	}
}
