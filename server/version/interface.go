package version

import (
	"github.com/ffenix113/gocraft/server/player"
	"github.com/ffenix113/gocraft/utils/types"
)

type Versioner interface {
	VersionName() string
	VersionNumber() int
	GamePacketHandlers() map[byte]PacketHandler
	ServerPacketHandlers() map[byte]PacketHandler
	RemoveGamePacketHandlers() []byte
	RemoveServerPacketHandlers() []byte
}

type PacketHandler func(pl *player.Player) (packets []types.Packet, err error)
