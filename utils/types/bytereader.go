package types

import "io"

type ByteReaderAddapter struct {
	r io.Reader
}

func (a ByteReaderAddapter) ReadByte() (byte, error) {
	b := make([]byte, 1)
	_, err := a.r.Read(b)
	return b[0], err
}
