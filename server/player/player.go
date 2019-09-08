package player

import (
	"net"
)

type State uint16

const (
	_ State = iota
	New
	HandshakeStart
	HandshakeEnd
	InGame
)

type Player struct {
	State State
	Conn  net.Conn
}
