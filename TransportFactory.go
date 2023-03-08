package goCommsDefinitions

type TransportFactory struct {
	Name       string
	StackNames []string
}

type TransportFactoryCallback func() *TransportFactory

func NewTransportFactory(name string, stackName ...string) TransportFactoryCallback {
	return func() *TransportFactory {
		return &TransportFactory{
			Name:       name,
			StackNames: stackName,
		}
	}
}
