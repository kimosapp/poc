package notification

type Service interface {
	SendWelcomeEmail(email string) error
	SendPasswordResetEmail(email string) error
	SendOrganizationInvitationEmail(email string) error
	SendOrganizationInvitationsEmail(email string) error
	SendOrganizationInvitationAcceptedEmail(email string) error
}
