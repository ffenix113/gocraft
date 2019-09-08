package main

import (
	"net"
	"os"

	"github.com/ffenix113/gocraft/server/player"

	"github.com/rs/zerolog"

	"github.com/ffenix113/gocraft/server"
	"github.com/ffenix113/gocraft/server/packer/plain"
)

func main() {
	serv := server.Server{
		Logger:         zerolog.New(os.Stdout).With().Str("component", "server").Logger(),
		TicksPerSecond: 1,
		Clients:        map[net.Conn]*player.Player{},
		PacketWriter: server.PacketWriter{
			Plain: plain.Uncompressed{},
		},
	}
	serv.Start()
}
