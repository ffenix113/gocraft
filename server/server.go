package server

import (
	"gocraft/server/packer"
	"gocraft/server/version"
	"gocraft/server/version/v404"
	"gocraft/utils/types"
	"log"
	"net"
	"runtime"

	"time"
)

const (
	maxPacketLength = 32767
)

type PacketWriter struct {
	Plain, Compressed packer.PacketWriter
}

type Server struct {
	Name           string
	TicksPerSecond int
	Version        version.Versioner
	PacketWriter   PacketWriter
	Clients        map[Player]struct{}
	serverPackets  map[byte]version.PacketHandler
	gamePackets    map[byte]version.PacketHandler
}

func (s *Server) Start() {
	s.serverPackets = map[byte]version.PacketHandler{}
	s.gamePackets = map[byte]version.PacketHandler{}
	s.Version = v404.NewVersion()
	for _, packet := range s.Version.ServerPacketHandlers() {
		s.serverPackets[packet.PackedId()] = packet.Handler
	}
	for _, packet := range s.Version.GamePacketHandlers() {
		s.gamePackets[packet.PackedId()] = packet.Handler
	}
	go s.AcceptNew()
	log.Println("server started")
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
		log.Println("new connection")
		go s.connectNewPlayer(conn)
	}

}

func (s *Server) Loop() {
	sleepTime := time.Second / time.Duration(s.TicksPerSecond)
	for {
		runtime.Gosched()
		time.Sleep(sleepTime)
	}
}

func (s *Server) connectNewPlayer(conn net.Conn) {
	var (
		//buf              = bytes.NewBuffer(make([]byte, 256))
		counter          uint8
		length, packetID types.VarInt
	)
	for {
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		counter++
		log.Println("request #", counter)

		//n, _ := conn.Read(buf.Bytes())
		//log.Printf("request: % X\n", buf.Bytes()[:n])

		length.Read(conn)
		packetID.Read(conn)
		log.Printf("length: %d, packetId: %X\n", length, packetID)
		if length.Value > maxPacketLength {
			log.Printf("packet length is %d, which is above maximum", length)
			conn.Close()
			return
		}

		if handler := s.serverPackets[byte(packetID.Value)]; length.Value != 0 && handler != nil {
			packets, err := handler(conn)
			if err != nil {
				log.Print(err)
				conn.Close()
				return
			}
			s.PacketWriter.Plain.Write(conn, packets)
			conn.Close()
			break
		} else {
			log.Printf("unknown server packet: %X", packetID)
		}
	}
}
