package client

import "github.com/kimosapp/poc/internal/core/ports/client"

type EmailClientGmail struct {
}

func NewEmailClientGmail() client.EmailClient {
	return &EmailClientGmail{}
}

func (client *EmailClientGmail) SendEmail(to, cc, cco []string, subject string, body string) error {
	return nil
}
