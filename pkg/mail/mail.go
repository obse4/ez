package mail

import "gopkg.in/gomail.v2"

type MailConfig struct {
	// 使用的邮箱
	From string
	// 邮箱校验
	AuthorizationCode string
	// host地址
	// 例如 smtp.qq.com
	Host string
	Mail *Mail
}

type Mail struct {
	from              string
	authorizationCode string
	host              string
}

type MailRequest struct {
	Title   string   `json:"title"`
	Message string   `json:"message"`
	MailTo  []string `json:"mail_to"`
}

func NewMail(p *MailConfig) *Mail {
	mail := Mail{
		from:              p.From,
		authorizationCode: p.AuthorizationCode,
		host:              p.Host,
	}
	p.Mail = &mail

	return p.Mail

}

func (s *Mail) Send(p MailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", p.MailTo...)
	m.SetHeader("Subject", p.Title)
	m.SetBody("text/html", p.Message)
	g := gomail.NewDialer(s.host, 465, s.from, s.authorizationCode)
	if err := g.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
