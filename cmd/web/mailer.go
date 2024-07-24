package main

import (
	"sync"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
	FromAddres string
	FromName   string
	Wait       *sync.WaitGroup
	MailerChan chan Message
	ErrorChan  chan error
	DoneChan   chan bool
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
	Template    string
}

// a function to listen for messages on the MailerChan
func (m *Mail) sendMail(msg Message, errChan chan error) {
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddres
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data

	// build html mail
	formattedMsg, err := m.buildHTMLMessage(msg)
	if err != nil {
		errChan <- err
	}

	// build plain mail
	plainMsg, err := m.buildPlainMessage(msg)
	if err != nil {
		errChan <- err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.encrypt(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smptClient, err := server.Connect()
	if err != nil {
		errChan <- err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)

	email.SetBody(mail.TextPlain, plainMsg)
	email.SetBody(mail.TextHTML, formattedMsg)

	if len(msg.Attachments) > 0 {
		for _, val := range msg.Attachments {
			email.AddAttachment(val)
		}
	}

	err = email.Send(smptClient)
	if err != nil {
		errChan <- err
	}
}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	return "", nil
}

func (m *Mail) buildPlainMessage(msg Message) (string, error) {
	return "", nil
}

func (m *Mail) encrypt(e string) mail.Encryption {
	switch e {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
