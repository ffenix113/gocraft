package server

import (
	"fmt"
	"log"
	"net"

	"github.com/ffenix113/gocraft/server/player"

	"github.com/rs/zerolog"

	"github.com/ffenix113/gocraft/server/packer"
	"github.com/ffenix113/gocraft/server/version"
	v404 "github.com/ffenix113/gocraft/server/version/v404"
	"github.com/ffenix113/gocraft/utils/types"

	"time"
)

const (
	maxPacketLength = 32767
)

type PacketWriter struct {
	Plain, Compressed packer.PacketWriter
}

type Server struct {
	Logger         zerolog.Logger
	Name           string
	TicksPerSecond int
	Version        version.Versioner
	PacketWriter   PacketWriter
	Clients        map[net.Conn]*player.Player
	serverPackets  map[byte]version.PacketHandler
	gamePackets    map[byte]version.PacketHandler
}

func (s *Server) Start() {
	s.serverPackets = map[byte]version.PacketHandler{}
	s.gamePackets = map[byte]version.PacketHandler{}
	s.Version = v404.NewVersion()
	for pID, packet := range s.Version.ServerPacketHandlers() {
		s.serverPackets[pID] = packet
	}
	for pID, packet := range s.Version.GamePacketHandlers() {
		s.gamePackets[pID] = packet
	}
	go s.AcceptNew()
	s.Logger.Info().Msg("server started")
	s.Loop()
}

func (s *Server) AcceptNew() {
	cn, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 25565,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := cn.Accept()
		go s.connectNewPlayer(conn)
	}

}

func (s *Server) Loop() {
	sleepTime := time.Second / time.Duration(s.TicksPerSecond)
	for {
		time.Sleep(sleepTime)
	}
}

func (s *Server) connectNewPlayer(conn net.Conn) {
	var (
		length, packetID types.VarInt
	)
	for {
		currPlayer, ok := s.Clients[conn]
		if !ok {
			currPlayer = &player.Player{
				State: player.New,
				Conn:  conn,
			}
			s.Clients[conn] = currPlayer
		}

		length.Read(conn)
		packetID.Read(conn)
		s.Logger.Info().Int32("length", length.Value).Str("packet_id", fmt.Sprintf("%X", packetID.Value)).Msg("")
		if length.Value == 0 || length.Value > maxPacketLength {
			s.Logger.Info().Int32("length", length.Value).Msg("connection closed, zero-length")
			delete(s.Clients, conn)
			conn.Close()
			break
		}

		var handler version.PacketHandler
		switch currPlayer.State {
		case player.InGame:
			handler = s.gamePackets[byte(packetID.Value)]
		default:
			handler = s.serverPackets[byte(packetID.Value)]
		}

		if length.Value == 0 || handler == nil {
			s.Logger.Info().Int32("packet_id", packetID.Value).Uint16("state", uint16(currPlayer.State)).Msg("unknown server packet")
			delete(s.Clients, conn)
			conn.Close()
			return
		}
		packets, err := handler(currPlayer)
		if err != nil {
			log.Print(err)
			delete(s.Clients, conn)
			conn.Close()
			return
		}
		s.PacketWriter.Plain.Write(conn, packets)
	}
}
