package types

import (
	"encoding/binary"
	"io"
)

type UShort struct {
	Value uint16
}

func (s *UShort) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &s.Value)
}

func (s *UShort) Write(w io.Writer) error {
	panic("me")
}
