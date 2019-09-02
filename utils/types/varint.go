package types

import (
	"encoding/binary"
	"io"
)

type VarInt struct {
	Value int32
}

func (i *VarInt) Read(r io.Reader) (err error) {
	vint, err := binary.ReadUvarint(ByteReaderAddapter{r})
	*i = VarInt{int32(uint32(vint))}
	return
}

func (i *VarInt) Write(w io.Writer) error {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, uint64(i.Value))
	_, err := w.Write(b[:n])
	return err
}

func WritePacketID(w io.Writer, packetId byte) error {
	return (&VarInt{int32(packetId)}).Write(w)
}
