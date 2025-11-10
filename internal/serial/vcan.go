package serial

import (
	"context"
	"net"

	"go.einride.tech/can/pkg/socketcan"
)

type VcanConnection struct {
	conn net.Conn
	Recv *socketcan.Receiver
}

func Connect(ctx context.Context, network, address string) (*VcanConnection, error) {
	connection, err := socketcan.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	return &VcanConnection{
		conn: connection,
		Recv: socketcan.NewReceiver(connection),
	}, nil
}

func (vc *VcanConnection) Close() error {
	if vc.conn == nil {
		return nil
	}

	return vc.conn.Close()
}
