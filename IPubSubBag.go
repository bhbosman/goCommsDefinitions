package goCommsDefinitions

type IAdder interface {
	Add(msg interface{})
}

type IPubSubBag interface {
	IAdder
	Close()
	Flush()
}
