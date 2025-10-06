package serial

import (
	"bufio"
	"io"

	"go.bug.st/serial"
)

type SerialReader struct {
	port   serial.Port
	Reader *bufio.Reader
}

func Open(portName string, baudRate int) (*SerialReader, error) {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}

	p, err := serial.Open(portName, mode)
	if err != nil {
		return nil, err
	}

	return &SerialReader{
		port:   p,
		Reader: bufio.NewReader(p),
	}, nil
}

func (s *SerialReader) ReadLine() (string, error) {
	return s.Reader.ReadString('\n')
}

func (s *SerialReader) Close() error {
	if s.port == nil {
		return nil
	}

	return s.port.Close()
}

var _ io.Closer = (*SerialReader)(nil)
