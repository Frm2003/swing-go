package protocol

import "encoding/binary"

type Message struct {
	ObjectID uint32
	OpCode   uint16
	Payload  []byte
}

func Decode(data []byte) *Message {
	return &Message{
		ObjectID: binary.LittleEndian.Uint32(data[0:4]),
		OpCode:   binary.LittleEndian.Uint16(data[4:6]),
		Payload:  data[8:],
	}
}

func Encode(message *Message) []byte {
	size := uint16(8 + len(message.Payload))

	buffer := make([]byte, size)

	binary.LittleEndian.PutUint32(buffer[0:4], uint32(message.ObjectID))
	binary.LittleEndian.PutUint16(buffer[4:6], uint16(message.OpCode))
	binary.LittleEndian.PutUint16(buffer[6:8], size)

	copy(buffer[8:], message.Payload)

	return buffer
}
