package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

const (
	maxReconnect = -1
)

func setOptions(ctx context.Context, hosts []string, nkey string, isTest bool) (nats.Options, error) {
	opts := nats.GetDefaultOptions()
	opts.Servers = hosts
	opts.Timeout = 2 * time.Second
	opts.RetryOnFailedConnect = true
	opts.MaxReconnect = maxReconnect

	opts.AsyncErrorCB = func(nc *nats.Conn, sub *nats.Subscription, err error) {
		log.Println("async error on subject:", sub.Subject, "status:", nc.Status(), "err:", err)
	}

	opts.ReconnectedCB = func(nc *nats.Conn) {
		log.Println("reconnected:", nc.ConnectedUrl())
	}

	opts.DisconnectedErrCB = func(nc *nats.Conn, err error) {
		log.Println("disconnected:", err)
	}

	opts.ClosedCB = func(nc *nats.Conn) {
		log.Println("connection closed:", nc.LastError())
	}

	if isTest {
		return opts, nil
	}

	// auth with nkey
	kp, err := nkeys.FromSeed([]byte(nkey))
	if err != nil {
		return nats.Options{}, fmt.Errorf("nkey error: %w", err)
	}

	pub, err := kp.PublicKey()
	if err != nil {
		return nats.Options{}, fmt.Errorf("pubkey error: %w", err)
	}

	opts.Nkey = pub
	opts.SignatureCB = func(nonce []byte) ([]byte, error) {
		return kp.Sign(nonce)
	}

	return opts, nil
}
