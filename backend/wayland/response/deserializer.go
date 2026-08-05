package response

import "encoding/binary"

type deSerializer struct {
	payload []byte
	offset  uint
}

func NewDeSerializer(payload []byte) *deSerializer {
	return &deSerializer{payload, 0}
}

func (d *deSerializer) Uint32() uint32 {
	v := binary.LittleEndian.Uint32(d.payload[d.offset : d.offset+4])
	d.offset += 4
	return v
}

func (d *deSerializer) String() string {
	size := d.Uint32()
	padd := (size + 3) &^ 3

	buf := d.payload[d.offset : d.offset+uint(size-1)]

	d.offset += uint(padd)

	return string(buf)
}
