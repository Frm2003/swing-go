package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

type Transport struct {
	conn *net.UnixConn
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

func connect() (*net.UnixConn, error) {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		return nil, fmt.Errorf("XDG_RUNTIME_DIR não definido")
	}

	display := os.Getenv("WAYLAND_DISPLAY")
	if display == "" {
		display = "wayland-0"
	}

	socketPath := filepath.Join(runtimeDir, display)

	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{
		Net:  "unix",
		Name: socketPath,
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (t *Transport) Send(message *Message) error {
	data := encode(message)

	var oob []byte
	if len(message.Fds) > 0 {
		oob = unix.UnixRights(message.Fds...)
	}

	total := 0

	for total < len(data) {
		n, _, err := t.conn.WriteMsgUnix(
			data[total:],
			oob,
			nil,
		)
		if err != nil {
			return err
		}

		total += n

		// Os ancillary data devem ser enviados junto com a primeira escrita.
		oob = nil
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
