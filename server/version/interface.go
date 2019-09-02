package version

import (
	"gocraft/utils/types"
	"io"
)

type Commander interface {
	PackedId() byte
	Data() []byte
	Handler(r io.Reader) (packets []types.Packet, err error)
}

type Versioner interface {
	VersionName() string
	VersionNumber() int
	GamePacketHandlers() []Commander
	ServerPacketHandlers() []Commander
	RemoveGamePacketHandlers() []byte
	RemoveServerPacketHandlers() []byte
}

type PacketHandler func(r io.Reader) (packets []types.Packet, err error)
