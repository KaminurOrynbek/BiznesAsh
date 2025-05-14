package _interface

type CombinedUsecase interface {
	NotificationUsecase
	VerificationUsecase
	SubscriptionUsecase
	EmailSender
	
	GetWelcomeEmailHTML() string
}
