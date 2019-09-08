package types

import (
	"io"
)

var (
	_           Typer = &VarLong{}
	varPartLong       = uint64(1<<7 - 1)
)

type VarLong struct {
	Value int64
}

func (l *VarLong) Read(r io.Reader) {
	var size uint
	var val uint64
	for {
		b, err := ByteReaderAdapter{r}.ReadByte()
		if err != nil {
			return
		}

		val |= (uint64(b) & varPartLong) << (size * 7)
		size++
		if size > 10 {
			return
		}

		if (b & 0x80) == 0 {
			break
		}
	}
	l.Value = int64(val)
}

func (l VarLong) Write(w io.Writer) {
	var b []byte
	ui := uint64(l.Value)
	for {
		if (ui & ^varPartLong) == 0 {
			b = append(b, byte(ui))
			break
		}
		b = append(b, byte((ui&varPartLong)|0x80))
		ui >>= 7
	}
	w.Write(b)
}
