package notification

import (
	"github.com/kimosapp/poc/internal/core/ports/client"
	notificationRepository "github.com/kimosapp/poc/internal/core/ports/repository/notifications"
	notificationsService "github.com/kimosapp/poc/internal/core/ports/service/notification"
)

type ServiceImpl struct {
	templateRepository     notificationRepository.NotificationTemplateRepository
	notificationRepository notificationRepository.NotificationRepository
	emailClient            *client.EmailClient
}

func NewNotificationService(

	emailClient *client.EmailClient,
) notificationsService.Service {
	return &ServiceImpl{emailClient: emailClient}
}

func (ns *ServiceImpl) SendWelcomeEmail(email string) error {

	return nil
}
func (ns *ServiceImpl) SendForgotPasswordEmail(email string, link string) error {
	return nil
}
func (ns *ServiceImpl) SendResetPasswordEmail(email string, date string) error {
	return nil
}
func (ns *ServiceImpl) SendOrganizationInvitationEmail(email string) error {
	return nil
}
func (ns *ServiceImpl) SendOrganizationInvitationsEmail(email string) error {
	return nil
}
func (ns *ServiceImpl) SendOrganizationInvitationAcceptedEmail(email string) error {
	return nil
}
