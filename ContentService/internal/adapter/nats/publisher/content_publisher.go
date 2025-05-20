package publisher

import (
	"encoding/json"
	"log"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/payloads"
	"github.com/KaminurOrynbek/BiznesAsh_lib/queue"
)

const (
	PostCreatedSubject    = "post.created"
	PostUpdatedSubject    = "post.updated"
	CommentCreatedSubject = "comment.created"
	PostReportedSubject   = "post.reported"
)

type ContentPublisher struct {
	queue queue.MessageQueue
}

func NewContentPublisher(q queue.MessageQueue) *ContentPublisher {
	return &ContentPublisher{queue: q}
}

func (p *ContentPublisher) publish(subject string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal event payload: %v", err)
		return err
	}
	return p.queue.Publish(subject, data)
}

func (p *ContentPublisher) PublishPostCreated(payload payloads.PostCreated) error {
	return p.publish(PostCreatedSubject, payload)
}

func (p *ContentPublisher) PublishPostUpdated(payload payloads.PostUpdated) error {
	return p.publish(PostUpdatedSubject, payload)
}

func (p *ContentPublisher) PublishCommentCreated(payload payloads.CommentCreated) error {
	return p.publish(CommentCreatedSubject, payload)
}

func (p *ContentPublisher) PublishPostReported(payload payloads.PostReported) error {
	return p.publish(PostReportedSubject, payload)
}
