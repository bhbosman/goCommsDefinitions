package goCommsDefinitions

import (
	"net"
	"time"
)

type ISpecificInformationForConnection interface {
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
