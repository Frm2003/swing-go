package backend

import (
	"io"
	"net"
)

type Transport struct {
	conn  net.Conn
	frame func(io.Reader) ([]byte, error)
}

type Connect func() (net.Conn, error)
type Frame func(io.Reader) ([]byte, error)

func NewTransport(connect Connect, frame Frame) *Transport {
	conn, err := connect()

	if err != nil {
		panic(err)
	}

	return &Transport{
		conn:  conn,
		frame: frame,
	}
}

func (t *Transport) Send(data []byte) error {
	var total = 0

	for total < len(data) {
		n, err := t.conn.Write(data[total:])

		if err != nil {
			return err
		}

		total += n
	}

	return nil
}

func (t *Transport) Receive() ([]byte, error) {
	return t.frame(t.conn)
}
