package request

import "encoding/binary"

type serializer struct {
	payload []byte
}

func NewSerializer() *serializer {
	return &serializer{}
}

func (s *serializer) Uint32(v uint32) *serializer {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, v)
	s.payload = append(s.payload, buf...)
	return s
}

func (s *serializer) String(v string) *serializer {
	size := len(v) + 1

	s.Uint32(uint32(size))

	padd := (size + 3) &^ 3
	buf := make([]byte, padd)

	copy(buf, v)

	s.payload = append(s.payload, buf...)

	return s
}

func (s *serializer) Bytes() []byte {
	return s.payload
}
