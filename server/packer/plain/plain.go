package plain

import (
	"bytes"
	"io"

	"github.com/ffenix113/gocraft/utils/types"
)

type Uncompressed struct{}

func (u Uncompressed) Write(w io.Writer, packets []types.Packet) error {
	var buf, outputBuf bytes.Buffer
	for _, packet := range packets {
		types.WritePacketID(&buf, packet.PacketID)
		for _, typ := range packet.Data {
			typ.Write(&buf)
		}
		(&types.VarInt{Value: int32(buf.Len())}).Write(&outputBuf)
		outputBuf.Write(buf.Bytes())
		//log.Printf("response: % X\n string: %s", outputBuf.Bytes(), outputBuf.String())
		w.Write(outputBuf.Bytes())
	}
	return nil
}
