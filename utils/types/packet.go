package types

import "io"

type Packet struct {
	PacketID byte
	Data     []Typer
}

func WritePacketID(w io.Writer, packetId byte) error {
	(&VarInt{Value: int32(packetId)}).Write(w)
	return nil
}
