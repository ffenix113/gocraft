package server

import (
	"context"
	"log"
	"net"
	"time"
)

type Player struct {
	conn net.Conn
}

func (p *Player) Communicate(ctx context.Context) {
	var b = make([]byte, 1)
	for {
		p.conn.SetDeadline(time.Now().Add(30 * time.Second))
		_, err := p.conn.Read(b)
		if err != nil {
			log.Println("read error:", err)
			p.conn.Close()
		}

	}
}
