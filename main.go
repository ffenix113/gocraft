package main

import (
	"gocraft/server"
	"gocraft/server/packer/plain"
)

func main() {
	serv := server.Server{
		TicksPerSecond: 1,
		PacketWriter: server.PacketWriter{
			Plain: plain.Plain{},
		},
	}
	serv.Start()
}
