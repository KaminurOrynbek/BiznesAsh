package subscriber

import (
	"encoding/json"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/payloads"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/KaminurOrynbek/BiznesAsh_lib/queue"
	"log"

	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type ContentSubscriber struct {
	queue               queue.MessageQueue
	notificationUsecase usecase.NotificationUsecase 
}

func NewContentSubscriber(q queue.MessageQueue, ns usecase.NotificationUsecase) *ContentSubscriber {
	return &ContentSubscriber{
		queue:               q,
		notificationUsecase: ns,
	}
}

func (s *ContentSubscriber) SubscribePostCreated() error {
	return s.queue.Subscribe("post.created", func(msg []byte) {
		var payload payloads.PostCreated
		err := json.Unmarshal(msg, &payload)
		if err != nil {
			log.Printf("Error unmarshalling PostCreated message: %v", err)
			return
		}
		s.handlePostCreated(payload)
	})
}

func (s *ContentSubscriber) SubscribePostUpdated() error {
	return s.queue.Subscribe("post.updated", func(msg []byte) {
		var payload payloads.PostUpdated
		err := json.Unmarshal(msg, &payload)
		if err != nil {
			log.Printf("Error unmarshalling PostUpdated message: %v", err)
			return
		}
		s.handlePostUpdated(payload)
	})
}

func (s *ContentSubscriber) SubscribeCommentCreated() error {
	return s.queue.Subscribe("comment.created", func(msg []byte) {
		var payload payloads.CommentCreated
		err := json.Unmarshal(msg, &payload)
		if err != nil {
			log.Printf("Error unmarshalling CommentCreated message: %v", err)
			return
		}
		s.handleCommentCreated(payload)
	})
}

func (s *ContentSubscriber) SubscribePostReported() error {
	return s.queue.Subscribe("post.reported", func(msg []byte) {
		var payload payloads.PostReported
		err := json.Unmarshal(msg, &payload)
		if err != nil {
			log.Printf("Error unmarshalling PostReported message: %v", err)
			return
		}
		s.handlePostReported(payload)
	})
}

func (s *ContentSubscriber) SubscribePostLiked() error {
	return s.queue.Subscribe("post.liked", func(msg []byte) {
		var payload payloads.PostLiked
		if err := json.Unmarshal(msg, &payload); err != nil {
			log.Printf("Failed to parse PostLiked: %v", err)
			return
		}
		s.handlePostLiked(payload)
	})
}

func (s *ContentSubscriber) SubscribeCommentLiked() error {
	return s.queue.Subscribe("comment.liked", func(msg []byte) {
		var payload payloads.CommentLiked
		if err := json.Unmarshal(msg, &payload); err != nil {
			log.Printf("Failed to parse CommentLiked: %v", err)
			return
		}
		s.handleCommentLiked(payload)
	})
}

func (s *ContentSubscriber) handlePostCreated(payload payloads.PostCreated) {
	log.Printf("New post created: %s by author: %s", payload.Title, payload.AuthorID)
	notification := &entity.Notification{
		UserID:  payload.AuthorID,
		Message: "A new post was created: " + payload.Title,
		PostID:  &payload.PostID,
	}
	err := s.notificationUsecase.NotifyNewPost(context.Background(), notification)
	if err != nil {
		log.Printf("Error sending new post notification: %v", err)
	}
}

func (s *ContentSubscriber) handlePostUpdated(payload payloads.PostUpdated) {
	log.Printf("Post updated: %s", payload.PostID)
	notification := &entity.Notification{
		UserID:  payload.AuthorID,
		Message: "Your post has been updated.",
		PostID:  &payload.PostID,
	}
	err := s.notificationUsecase.NotifyPostUpdate(context.Background(), notification)
	if err != nil {
		log.Printf("Error sending post update notification: %v", err)
	}
}

func (s *ContentSubscriber) handleCommentCreated(payload payloads.CommentCreated) {
	log.Printf("New comment created on post: %s by user: %s", payload.PostID, payload.UserID)
	notification := &entity.Notification{
		UserID:    payload.UserID,
		Message:   "New comment created on your post.",
		PostID:    &payload.PostID,
		CommentID: &payload.CommentID,
	}
	err := s.notificationUsecase.SendCommentNotification(context.Background(), notification)
	if err != nil {
		log.Printf("Error sending comment notification: %v", err)
	}
}

func (s *ContentSubscriber) handlePostReported(payload payloads.PostReported) {
	log.Printf("Post reported: %s by reporter: %s, reason: %s", payload.PostID, payload.ReporterID, payload.Reason)
	notification := &entity.Notification{
		UserID:  payload.ReporterID,
		Message: "A post has been reported: " + payload.PostID,
		PostID:  &payload.PostID,
	}
	err := s.notificationUsecase.SendReportNotification(context.Background(), notification)
	if err != nil {
		log.Printf("Error sending report notification: %v", err)
	}
}

func (s *ContentSubscriber) handlePostLiked(payload payloads.PostLiked) {
	log.Printf("Post liked: %s (notify user %s)", payload.PostID, payload.UserID)
	notification := &entity.Notification{
		UserID:  payload.UserID,
		Message: "Your post got a new like!",
		PostID:  &payload.PostID,
	}
	err := s.notificationUsecase.NotifyPostLike(context.Background(), notification)
	if err != nil {
		log.Printf("Failed to notify post like: %v", err)
	}
}

func (s *ContentSubscriber) handleCommentLiked(payload payloads.CommentLiked) {
	log.Printf("Comment liked: %s (notify user %s)", payload.CommentID, payload.UserID)
	notification := &entity.Notification{
		UserID:    payload.UserID,
		Message:   "Your comment got a new like!",
		CommentID: &payload.CommentID,
	}
	err := s.notificationUsecase.NotifyCommentLike(context.Background(), notification)
	if err != nil {
		log.Printf("Failed to notify comment like: %v", err)
	}
}
