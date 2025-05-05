package queue

type MessageQueue interface {
	Publish(subject string, message []byte) error
	Subscribe(subject string, handler func(msg []byte)) error
	Close() error
}
