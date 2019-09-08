package types

import (
	"io"
)

var _ Typer = &String{}

type String struct {
	Value string
}

func (s *String) Read(r io.Reader) {
	var length VarInt
	length.Read(r)
	bts := make([]byte, length.Value)
	r.Read(bts)
	s.Value = string(bts)
}
func (s *String) Write(w io.Writer) {
	v := VarInt{int32(len(s.Value))}
	v.Write(w)
	w.Write([]byte(s.Value))
}
