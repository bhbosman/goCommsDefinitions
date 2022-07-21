package goCommsDefinitions

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
