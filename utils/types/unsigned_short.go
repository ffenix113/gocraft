package types

import (
	"encoding/binary"
	"io"
)

var _ Typer = &UShort{}

type UShort struct {
	Value uint16
}

func (s *UShort) Read(r io.Reader) {
	binary.Read(r, binary.BigEndian, &s.Value)
}

func (s *UShort) Write(w io.Writer) {
	binary.Write(w, binary.BigEndian, s.Value)
}
