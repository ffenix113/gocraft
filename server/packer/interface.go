package packer

import (
	"github.com/ffenix113/gocraft/utils/types"
	"io"
)

type PacketWriter interface {
	Write(w io.Writer, packets []types.Packet) error
}
