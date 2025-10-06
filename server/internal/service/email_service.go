package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/lubosgarancovsky/eden-arc/internal/config"
)

type EmailService struct {
	cfg *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

func (s *EmailService) SendTemplateEmail(to, subject, templatePath string, data any) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	msg := []byte(fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("From: %s\r\n", s.cfg.SMTPFrom) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body.String())

	addr := fmt.Sprintf("%s:%d", s.cfg.SMTPHost, s.cfg.SMTPPort)

	if err := smtp.SendMail(addr, nil, s.cfg.SMTPFrom, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
