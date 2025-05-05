package publisher

import (
	"encoding/json"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/pkg/queue"
	"log"
)

const (
	UserRegisteredSubject          = "user.registered"
	UserDeletedSubject             = "user.deleted"
	UserPromotedToModeratorSubject = "user.promoted_to_moderator"
	UserPromotedToAdminSubject     = "user.promoted_to_admin"
	UserDemotedSubject             = "user.demoted"
	UserBannedSubject              = "user.banned"
	ReportCreatedSubject           = "report.created"
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

func (p *UserPublisher) PublishUserRegistered(userId, email string) error {
	return p.publish(UserRegisteredSubject, map[string]string{
		"user_id": userId,
		"email":   email,
	})
}

func (p *UserPublisher) PublishUserDeleted(userId string) error {
	return p.publish(UserDeletedSubject, map[string]string{
		"user_id": userId,
	})
}

func (p *UserPublisher) PublishUserPromotedToModerator(userId string) error {
	return p.publish(UserPromotedToModeratorSubject, map[string]string{
		"user_id": userId,
	})
}

func (p *UserPublisher) PublishUserPromotedToAdmin(userId string) error {
	return p.publish(UserPromotedToAdminSubject, map[string]string{
		"user_id": userId,
	})
}

func (p *UserPublisher) PublishUserDemoted(userId string) error {
	return p.publish(UserDemotedSubject, map[string]string{
		"user_id": userId,
	})
}

func (p *UserPublisher) PublishUserBanned(userId string, reason string) error {
	return p.publish(UserBannedSubject, map[string]string{
		"user_id": userId,
		"reason":  reason,
	})
}

func (p *UserPublisher) PublishReportCreated(reportId string) error {
	return p.publish(ReportCreatedSubject, map[string]string{
		"report_id": reportId,
	})
}
