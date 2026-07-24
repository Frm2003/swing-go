package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func Connect() (net.Conn, error) {
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

func Frame(conn io.Reader) ([]byte, error) {
	header := make([]byte, 8)

	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	size := binary.LittleEndian.Uint16(header[6:8])
	payload := make([]byte, size-8)

	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	result := make([]byte, 8+len(payload))

	copy(result[:8], header)
	copy(result[8:], payload)

	return result, nil
}
