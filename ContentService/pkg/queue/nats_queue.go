package queue

import (
	"github.com/nats-io/nats.go"
	"log"
)

type NATSQueue struct {
	conn *nats.Conn
}

func NewNATSQueue(conn *nats.Conn) *NATSQueue {
	return &NATSQueue{conn: conn}
}

func (n *NATSQueue) Publish(subject string, message []byte) error {
	return n.conn.Publish(subject, message)
}

func (n *NATSQueue) Subscribe(subject string, handler func(msg []byte)) error {
	_, err := n.conn.Subscribe(subject, func(m *nats.Msg) {
		handler(m.Data)
	})
	return err
}

func (n *NATSQueue) Close() error {
	n.conn.Close()
	log.Println("NATS connection closed")
	return nil
}
