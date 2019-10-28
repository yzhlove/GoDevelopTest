package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Packet struct {
	Length uint16
	Body   []byte
}

func UnPacket(reader io.Reader) (pkg Packet, err error) {

	var buffer = make([]byte, 2)
	if _, err = io.ReadFull(reader, buffer); err != nil {
		fmt.Println("ReadFullErr:", err)
		return
	}
	newReader := bytes.NewReader(buffer)
	if err = binary.Read(newReader, binary.LittleEndian, &pkg.Length); err != nil {
		fmt.Println("binaryReadErr:", err)
		return
	}
	pkg.Body = make([]byte, pkg.Length)
	_, err = io.ReadFull(reader, pkg.Body)
	return
}

func Pack(reader io.Writer, data []byte) (err error) {
	var buffer bytes.Buffer
	if err = binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		fmt.Println("binaryWriteErr:", err)
		return
	}
	if _, err = buffer.Write(data); err != nil {
		fmt.Println("bufferWriteErr:", err)
		return
	}
	out := buffer.Bytes()
	_, err = reader.Write(out)
	return
}
