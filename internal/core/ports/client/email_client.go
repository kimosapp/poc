package client

type EmailClient interface {
	SendEmail(to, cc, cco []string, subject string, body string) error
}
