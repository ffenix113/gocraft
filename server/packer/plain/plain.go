package plain

import (
	"bytes"
	"gocraft/utils/types"
	"io"
	"log"
)

type Plain struct {
}

func (p Plain) Write(w io.Writer, packets []types.Packet) error {
	var buf, outputBuf bytes.Buffer
	for _, packet := range packets {
		for _, typ := range packet.Data {
			typ.Write(&buf)
		}
		(&types.VarInt{int32(buf.Len())}).Write(&outputBuf)
		types.WritePacketID(&outputBuf, packet.PacketID)
		outputBuf.Write(buf.Bytes())
		log.Printf("responce: % X\n string: %s", outputBuf.Bytes(), outputBuf.String())
		w.Write(outputBuf.Bytes())
	}
	return nil
}
