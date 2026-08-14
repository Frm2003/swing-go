package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

type Transport struct {
	conn net.Conn
}

func NewTransport() *Transport {
	conn, err := connect()

	if err != nil {
		panic(err)
	}

	return &Transport{
		conn: conn,
	}
}

func connect() (net.Conn, error) {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")

	if runtimeDir == "" {
		return nil, fmt.Errorf("XDG_RUNTIME_DIR não definido")
	}

	display := os.Getenv("WAYLAND_DISPLAY")

	if display == "" {
		display = "wayland-0"
	}

	socketPath := filepath.Join(runtimeDir, display)

	return net.Dial("unix", socketPath)
}

func (t *Transport) Send(message *Message) error {
	var total = 0

	data := encode(message)

	for total < len(data) {
		n, err := t.conn.Write(data[total:])

		if err != nil {
			return err
		}

		total += n
	}

	return nil
}

func (t *Transport) Receive() (*Message, error) {
	header := make([]byte, 8)

	if _, err := io.ReadFull(t.conn, header); err != nil {
		return nil, err
	}

	size := binary.LittleEndian.Uint16(header[6:8])
	payload := make([]byte, size-8)

	if _, err := io.ReadFull(t.conn, payload); err != nil {
		return nil, err
	}

	result := make([]byte, 8+len(payload))

	copy(result[:8], header)
	copy(result[8:], payload)

	return decode(result), nil
}
