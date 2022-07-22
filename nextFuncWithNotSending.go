package goCommsDefinitions

import "github.com/reactivex/rxgo/v2"

type TryNextFunc func(interface{}) bool
type IsNextActive func() bool

func CreateTryNextFunc(channel chan<- interface{}) TryNextFunc {
	errHappen := false
	return func(data interface{}) bool {
		defer func() {
			if err := recover(); err != nil {
				errHappen = true
			}
		}()
		if errHappen {
			return false
		}
		select {
		case channel <- data:
			return true
		default:
			return false
		}
	}
}

func CreateNextFunc(channel chan<- interface{}) rxgo.NextFunc {
	errHappen := false
	return func(data interface{}) {
		defer func() {
			if err := recover(); err != nil {
				errHappen = true
			}
		}()
		if errHappen {
			return
		}
		channel <- data
	}
}
