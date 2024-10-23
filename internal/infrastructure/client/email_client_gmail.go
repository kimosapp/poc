package client

import (
	"github.com/kimosapp/poc/internal/core/ports/client"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	"github.com/kimosapp/poc/internal/infrastructure/configuration"
	mail "github.com/xhit/go-simple-mail/v2"
)

type EmailClientGmail struct {
	server *mail.SMTPServer
	from   string
	log    logging.Logger
}

func NewEmailClientGmail(
	log logging.Logger,
) client.EmailClient {
	config := configuration.GetGmailClientConfig()
	server := mail.NewSMTPClient()
	server.Host = config.GetUrl()
	server.Port = config.GetPort()
	server.Username = config.GetUser()
	server.Password = config.GetPassword()
	server.Encryption = mail.EncryptionTLS
	_, err := server.Connect()
	if err != nil {
		log.Fatal("Error connecting to smtp server", "error", err)
	}
	return &EmailClientGmail{server: server, log: log, from: config.GetUser()}

}

func (c *EmailClientGmail) SendEmail(to, cc, cco []string, subject string, body string) error {
	smtpClient, err := c.server.Connect()
	if err != nil {
		//TODO add something here to execute a fallback
		c.log.Fatal("Error connecting to smtp server", "error", err)
	}
	email := mail.NewMSG()
	email.SetFrom(c.from)
	email.AddTo(to...)
	email.AddCc(cc...)
	email.AddBcc(cco...)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	err = email.Send(smtpClient)
	return err
}
