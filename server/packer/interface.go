package packer

import (
	"gocraft/utils/types"
	"io"
)

type PacketWriter interface {
	Write(w io.Writer, packets []types.Packet) error
}
