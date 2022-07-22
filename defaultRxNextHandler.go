package goCommsDefinitions

import (
	"github.com/reactivex/rxgo/v2"
)

type IRxNextHandler interface {
	OnSendData(i interface{})
	OnError(err error)
	OnComplete()
	OnTrySendData(i interface{}) bool
	IsActive() bool
}

type DefaultRxNextHandler struct {
	i      rxgo.NextFunc
	e      rxgo.ErrFunc
	c      rxgo.CompletedFunc
	t      TryNextFunc
	active IsNextActive
}

func NewDefaultRxNextHandler(
	i rxgo.NextFunc,
	t TryNextFunc,
	e rxgo.ErrFunc,
	c rxgo.CompletedFunc,
	active IsNextActive,
) *DefaultRxNextHandler {
	return &DefaultRxNextHandler{
		i: func(i rxgo.NextFunc) rxgo.NextFunc {
			if i != nil {
				return i
			}
			return func(i interface{}) {

			}
		}(i),
		t: func(t TryNextFunc) TryNextFunc {
			if t != nil {
				return t
			}
			return func(i interface{}) bool {
				return false
			}
		}(t),
		e: e,
		c: c,
		active: func(isActive func() bool) func() bool {
			if isActive != nil {
				return isActive
			}
			return func() bool {
				return true
			}
		}(active),
	}
}

func (self *DefaultRxNextHandler) OnSendData(i interface{}) {
	self.i(i)
}

func (self *DefaultRxNextHandler) OnTrySendData(i interface{}) bool {
	return self.t(i)
}

func (self *DefaultRxNextHandler) OnError(err error) {
	if self.e != nil {
		self.e(err)
	}
}

func (self *DefaultRxNextHandler) OnComplete() {
	if self.c != nil {
		self.c()
	}
}

func (self *DefaultRxNextHandler) IsActive() bool {
	return self.active()
}
