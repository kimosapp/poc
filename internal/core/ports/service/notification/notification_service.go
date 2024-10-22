package notification

type Service interface {
	SendWelcomeEmail(email string) error
	SendForgotPasswordEmail(email string, link string) error
	SendResetPasswordEmail(email string, date string) error
	SendOrganizationInvitationEmail(email string) error
	SendOrganizationInvitationsEmail(email string) error
	SendOrganizationInvitationAcceptedEmail(email string) error
}
