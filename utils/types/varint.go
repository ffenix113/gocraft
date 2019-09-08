package types

import (
	"encoding/binary"
	"io"
)

var (
	_       Typer = &VarInt{}
	varPart       = uint32(0x7F)
)

type VarInt struct {
	Value int32
}

func (i *VarInt) Read(r io.Reader) {
	vint, _ := binary.ReadUvarint(ByteReaderAdapter{r})
	*i = VarInt{int32(vint)}
}

func (i *VarInt) Write(w io.Writer) {
	var b []byte
	ui := uint32(i.Value)
	for {
		if (ui & ^varPart) == 0 {
			b = append(b, byte(ui))
			break
		}
		b = append(b, byte((ui&varPart)|0x80))
		ui >>= 7
	}
	w.Write(b)
}
