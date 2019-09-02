package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gocraft/utils/types"
	"io"
	"log"
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

func (h Handshake) Data() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, h.protoVersion)
	binary.Write(&buf, binary.BigEndian, h.address)
	binary.Write(&buf, binary.BigEndian, h.port)
	binary.Write(&buf, binary.BigEndian, h.nextState)
	return buf.Bytes()
}

func (h Handshake) Handler(r io.Reader) ([]types.Packet, error) {
	h.protoVersion.Read(r)
	log.Println("read proto:", h.protoVersion.Value)
	h.address.Read(r)
	log.Println("read address:", h.address.Value)
	h.port.Read(r)
	log.Println("read port:", h.port.Value)
	h.nextState.Read(r)
	log.Println("read next state:", h.nextState.Value)
	log.Printf("got data from handshake: %v\n", h)

	b := make([]byte, 2)
	r.Read(b)
	if b[0] != 0x01 && b[1] != 0x00 {
		log.Println("invalid next sequence: ", b)
		return nil, errors.New("invalid handshake")
	}

	if h.nextState.Value == 2 {
		return nil, errors.New("not implemented")
	}
	return []types.Packet{
		{
			PacketID: packetId,
			Data: []types.Typer{
				&types.String{`{
    "version": {
        "name": "1.13.2",
        "protocol": 404
    },
    "players": {
        "max": 100,
        "online": 5,
        "sample": [
            {
                "name": "thinkofdeath",
                "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
            }
        ]
    },	
    "description": {
        "text": "Hello world"
    },
    "favicon": ""
}`},
			},
		},
	}, nil
}
