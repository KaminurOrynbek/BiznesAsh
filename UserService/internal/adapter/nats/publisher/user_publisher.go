package publisher

import (
	"encoding/json"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/nats/payloads"
	"github.com/KaminurOrynbek/BiznesAsh_lib/queue"
	"log"
)

const (
	UserRegisteredSubject          = "user.registered"
	UserDeletedSubject             = "user.deleted"
	UserPromotedToModeratorSubject = "user.promoted_to_moderator"
	UserPromotedToAdminSubject     = "user.promoted_to_admin"
	UserDemotedSubject             = "user.demoted"
	UserBannedSubject              = "user.banned"
)

type UserPublisher struct {
	queue queue.MessageQueue
}

func NewUserPublisher(q queue.MessageQueue) *UserPublisher {
	return &UserPublisher{queue: q}
}

func (p *UserPublisher) publish(subject string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal event payload: %v", err)
		return err
	}
	return p.queue.Publish(subject, data)
}

func (p *UserPublisher) PublishUserRegistered(payload payloads.UserEventPayload) error {
	return p.publish(UserRegisteredSubject, payload)
}

func (p *UserPublisher) PublishUserDeleted(payload payloads.UserEventPayload) error {
	return p.publish(UserDeletedSubject, payload)
}

func (p *UserPublisher) PublishUserPromotedToModerator(payload payloads.UserEventPayload) error {
	return p.publish(UserPromotedToModeratorSubject, payload)
}

func (p *UserPublisher) PublishUserPromotedToAdmin(payload payloads.UserEventPayload) error {
	return p.publish(UserPromotedToAdminSubject, payload)
}

func (p *UserPublisher) PublishUserDemoted(payload payloads.UserEventPayload) error {
	return p.publish(UserDemotedSubject, payload)
}

func (p *UserPublisher) PublishUserBanned(payload payloads.UserEventPayload) error {
	return p.publish(UserBannedSubject, payload)
}
