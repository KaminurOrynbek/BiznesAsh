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
	PostLikedSubject      = "post.liked"
	CommentLikedSubject   = "comment.liked"
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
		log.Printf("[PUBLISH ERROR] Failed to marshal payload for subject %s: %v", subject, err)
		return err
	}

	log.Printf("[PUBLISH] Subject: %s | Payload: %s", subject, string(data))

	err = p.queue.Publish(subject, data)
	if err != nil {
		log.Printf("[PUBLISH ERROR] Failed to publish to subject %s: %v", subject, err)
	}
	return err
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

func (p *ContentPublisher) PublishPostLiked(payload payloads.PostLiked) error {
	return p.publish(PostLikedSubject, payload)
}

func (p *ContentPublisher) PublishCommentLiked(payload payloads.CommentLiked) error {
	return p.publish(CommentLikedSubject, payload)
}
