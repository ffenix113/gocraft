package types

import (
	"io"
)

type String struct {
	Value string
}

func (s *String) Read(r io.Reader) error {
	var length VarInt
	length.Read(r)
	bts := make([]byte, length.Value)
	r.Read(bts)
	s.Value = string(bts)
	return nil
}
func (s *String) Write(w io.Writer) error {
	v := VarInt{int32(len(s.Value))}
	v.Write(w)
	w.Write([]byte(s.Value))
	return nil
}
