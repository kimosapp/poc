package configuration

import (
	"os"
	"strconv"
)

type GmailClientConfig struct {
	user     string
	password string
	url      string
	port     int
}

func (c *GmailClientConfig) GetPort() int {
	return c.port
}
func (c *GmailClientConfig) GetPassword() string {
	return c.password
}
func (c *GmailClientConfig) GetUrl() string {
	return c.url
}
func (c *GmailClientConfig) GetUser() string {
	return c.user
}

var gmailClientConfig *GmailClientConfig

func GetGmailClientConfig() *GmailClientConfig {
	if gmailClientConfig == nil {
		initGmailClientConfig()
	}
	return gmailClientConfig
}

func initGmailClientConfig() {
	var err error
	gmailClientConfig = &GmailClientConfig{}
	gmailClientConfig.user = os.Getenv("GMAIL_EMAIL_FROM")
	gmailClientConfig.password = os.Getenv("GMAIL_EMAIL_PASS")
	gmailClientConfig.url = os.Getenv("GMAIL_SMTP_HOST")
	gmailClientConfig.port, err = strconv.Atoi(os.Getenv("GMAIL_SMTP_PORT"))
	if err != nil {
		panic("Error parsing the gmail port")
	}
}
