package version

type Basic struct {
	versionName   string
	versionNumber int
	gamePackets   []Commander
	serverPackets []Commander
}

func NewBasic(name string, number int) *Basic {
	return &Basic{
		versionName:   name,
		versionNumber: number,
		gamePackets:   []Commander{},
		serverPackets: []Commander{},
	}
}

func (b Basic) VersionName() string {
	return b.versionName
}

func (b Basic) VersionNumber() int {
	return b.versionNumber
}

func (b Basic) GamePacketHandlers() []Commander {
	return b.gamePackets
}

func (b Basic) ServerPacketHandlers() []Commander {
	return b.serverPackets
}

func (b Basic) RemoveGamePacketHandlers() []byte {
	return nil
}

func (b Basic) RemoveServerPacketHandlers() []byte {
	return nil
}
