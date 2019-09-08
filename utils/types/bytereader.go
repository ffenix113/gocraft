package types

import "io"

type ByteReaderAdapter struct {
	r io.Reader
}

func (a ByteReaderAdapter) ReadByte() (byte, error) {
	b := make([]byte, 1)
	_, err := a.r.Read(b)
	return b[0], err
}
