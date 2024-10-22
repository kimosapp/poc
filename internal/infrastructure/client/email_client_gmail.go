package client

import (
	"github.com/kimosapp/poc/internal/core/ports/client"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
)

type EmailClientGmail struct {
	server *mail.SMTPServer
	from   string
	log    logging.Logger
}

func NewEmailClientGmail(
	from string,
	pass string,
	smtpHost string,
	smtpPort string,
	log logging.Logger,
) client.EmailClient {
	server := mail.NewSMTPClient()
	server.Host = smtpHost
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Error("Error converting port to int", "error", err)
		panic(err)
	}
	server.Port = port
	server.Username = from
	server.Password = pass
	server.Encryption = mail.EncryptionTLS
	_, err = server.Connect()
	if err != nil {
		log.Fatal("Error connecting to smtp server", "error", err)
	}
	return &EmailClientGmail{server: server, log: log, from: from}

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
