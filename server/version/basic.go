package version

type Basic struct {
	versionName   string
	versionNumber int
	gamePackets   map[byte]PacketHandler
	serverPackets map[byte]PacketHandler
}

func NewBasic(name string, number int) *Basic {
	return &Basic{
		versionName:   name,
		versionNumber: number,
	}
}

func (b Basic) VersionName() string {
	return b.versionName
}

func (b Basic) VersionNumber() int {
	return b.versionNumber
}

func (b Basic) GamePacketHandlers() map[byte]PacketHandler {
	return b.gamePackets
}

func (b Basic) ServerPacketHandlers() map[byte]PacketHandler {
	return b.serverPackets
}

func (b Basic) RemoveGamePacketHandlers() []byte {
	return nil
}

func (b Basic) RemoveServerPacketHandlers() []byte {
	return nil
}
