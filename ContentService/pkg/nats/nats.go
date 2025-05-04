package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
)

type Client struct {
	Conn *nats.Conn
}

// NewClient connects to the NATS server using options from setOptions.
func NewClient(ctx context.Context, hosts []string, nkey string, isTest bool) (*Client, error) {
	opts, err := setOptions(ctx, hosts, nkey, isTest)
	if err != nil {
		return nil, fmt.Errorf("setOptions: %w", err)
	}

	conn, err := opts.Connect()
	if err != nil {
		return nil, fmt.Errorf("opts.Connect: %w", err)
	}

	return &Client{Conn: conn}, nil
}

func (nc *Client) Close() {
	nc.Conn.Close()
}
