package server

import (
	"errors"
	"log"

	"github.com/ffenix113/gocraft/server/player"

	"github.com/ffenix113/gocraft/utils/types"
)

const packetId = 0x00

type Handshake struct {
	protoVersion types.VarInt
	address      types.String
	port         types.UShort
	nextState    types.VarInt
}

func (h Handshake) PackedId() byte {
	return packetId
}

func HandshakeHandler(pl *player.Player) ([]types.Packet, error) {
	var h Handshake
	h.protoVersion.Read(pl.Conn)
	h.address.Read(pl.Conn)
	h.port.Read(pl.Conn)
	h.nextState.Read(pl.Conn)
	log.Printf("got data from handshake: %#v\n", h)

	b := make([]byte, 2)
	pl.Conn.Read(b)
	if b[0] != 0x01 && b[1] != 0x00 {
		log.Println("invalid next sequence: ", b)
		return nil, errors.New("invalid handshake")
	}

	switch {
	case h.nextState.Value == 1 && pl.State == player.New:
		return ServerState(nil)
	case h.nextState.Value == 1:
		return PongHandshake(pl)
	case h.nextState.Value == 2:
		return LoginHandler(pl)
	}
	panic("should not be here")
}

func ServerState(_ *player.Player) ([]types.Packet, error) {
	return []types.Packet{
		{
			PacketID: packetId,
			Data: []types.Typer{
				&types.String{Value: `{
    "version": {
        "name": "1.13.2",
        "protocol": 404
    },
    "players": {
        "max": 4,
        "online": 1,
        "sample": [
            {
                "name": "thinkofdeath",
                "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
            }
        ]
    },	
    "description": {
        "text": "Hello from Golang!"
    },
    "favicon": ""
}`},
			},
		},
	}, nil
}

func PongHandshake(pl *player.Player) ([]types.Packet, error) {
	pl.Conn.Read(make([]byte, 1))
	var l types.VarLong
	l.Read(pl.Conn)

	return []types.Packet{
		{
			PacketID: 0x01,
			Data: []types.Typer{
				&l,
			},
		},
	}, nil
}
