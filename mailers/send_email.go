package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

// SendEmail sends an email to an email id
func SendEmail(email, content string) error {
	m := mail.NewMessage()

	m.Subject = "Event invitation!"
	m.From = "karuppiah7890@gmail.com"
	m.To = []string{email}
	m.Subject = "Event invitation!"
	err := m.AddBody(r.String(content), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
