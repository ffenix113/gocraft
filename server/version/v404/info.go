package v404

import (
	"gocraft/server/version"
	"gocraft/server/version/v404/server"
)

type Version struct {
	*version.Basic
}

func NewVersion() version.Versioner {
	return &Version{
		Basic: version.NewBasic("1.13.2", 404),
	}
}

func (v *Version) ServerPacketHandlers() []version.Commander {
	return []version.Commander{
		server.Handshake{},
	}
}
