package server

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
)

// Mail Mail is a mail micro service
// The mail micro service could be using as a admin mailer,
// which could send mail to users.
type Mail struct {
	log Logger

	username     string
	password     string
	host         string
	address      string
	from         string
	tempaltePath string

	auth     smtp.Auth
	template *template.Template
}

// NewMail NewMail returns a mail micro service with account config.
//
// The username and password is used to authenticate to host.
//
// The address must include a port, as in "mail.example.com:smtp".
//
// The tempaltePath provide template for sending a html mail.
func NewMail(
	log Logger,
	username,
	password,
	host,
	address,
	from,
	templatePath string,
) *Mail {
	m := &Mail{
		log,
		username,
		password,
		host,
		address,
		from,
		templatePath,
		nil,
		nil,
	}

	m.auth = smtp.PlainAuth("", m.username, m.password, m.host)

	if tpl, err := template.ParseGlob(templatePath); err != nil {
		m.log.Printf("[Error] load tempalte from %s failed, cause of: %v\n", templatePath, err)
		m.template = template.New("default")
	} else {
		m.template = tpl
	}
	return m
}

// Send Send a mail to "to", whose body is the template render with data.
func (m *Mail) Send(to []string, subject, templateName string, data interface{}) error {
	var buf bytes.Buffer
	if err := m.template.ExecuteTemplate(&buf, templateName, data); err != nil {
		m.log.Printf(
			"[Warn] rend template %s failed when send mail %s, try to write data as []byte",
			templateName,
			subject,
		)
		if body, ok := data.([]byte); ok {
			buf.Write(body)
		}
	}
	msg := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))),
		[]byte(fmt.Sprintf("Subject: %s\r\n", subject)),
		[]byte("\r\n"),
		buf.Bytes(),
		[]byte("\r\n"),
	}, []byte(""))

	err := smtp.SendMail(m.address, m.auth, m.from, to, msg)
	if err != nil {
		m.log.Printf("[Error] send mail %s to %v failed, cause of: %v\n", subject, to, err)
		return errSendMail
	}
	return nil
}
