package server

import (
	"github.com/ffenix113/gocraft/server/player"
	"github.com/ffenix113/gocraft/utils/types"
)

type Login struct {
	Username types.String
}

func LoginHandler(pl *player.Player) (packets []types.Packet, err error) {
	var l Login
	l.Username.Read(pl.Conn)
	pl.State = player.InGame

	return []types.Packet{
		{
			PacketID: 0x02,
			Data: []types.Typer{
				types.NewUUIDFromUsername(l.Username.Value, true).AsString(),
				&l.Username,
			},
		},
	}, nil
}
