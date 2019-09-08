package v404

import (
	"github.com/ffenix113/gocraft/server/version"
	"github.com/ffenix113/gocraft/server/version/v404/server"
)

type Version struct {
	*version.Basic
}

func NewVersion() version.Versioner {
	return &Version{
		Basic: version.NewBasic("1.13.2", 404),
	}
}

func (v *Version) ServerPacketHandlers() map[byte]version.PacketHandler {
	return map[byte]version.PacketHandler{
		0x00: server.HandshakeHandler,
		0x01: server.PongHandshake,
	}
}
