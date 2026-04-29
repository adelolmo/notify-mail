package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"os"
	"path/filepath"
)

type Mail struct {
	Sender         string
	Authentication smtp.Auth
}

func NewNotification(account, password string) (*Mail, error) {
	if account == "" {
		return nil, fmt.Errorf("account is required")
	}
	if password == "" {
		return nil, fmt.Errorf("password is required")
	}
	auth := smtp.PlainAuth(
		"Notification",
		account,
		password,
		"smtp.gmail.com",
	)
	return &Mail{
		Sender:         "Notification",
		Authentication: auth,
	}, nil
}

func (m *Mail) Send(recipient, subject, message string, attachments ...string) error {
	return send(m.Authentication, m.Sender, recipient, subject, message, attachments)
}

func (m *Mail) SendTemplate(recipient, subject, templateFilename string, variables map[string]string, attachments ...string) error {
	template := NewTemplate(variables)
	message, err := template.ReplaceContent(templateFilename)
	if err != nil {
		return fmt.Errorf("cannot replace variables for placeholders in template file %s. Error: %s",
			templateFilename, err)
	}

	return send(m.Authentication, m.Sender, recipient, subject, message, attachments)
}

func send(authentication smtp.Auth, sender, recipient, subject, message string, attachments []string) error {
	if err := validateAttachments(attachments); err != nil {
		return err
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	header := map[string]string{
		"From":         sender,
		"To":           recipient,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": fmt.Sprintf("multipart/mixed;\r\n\tboundary=\"%s\"", writer.Boundary()),
	}
	for k, v := range header {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	buf.WriteString("\r\n")

	part, err := writer.CreatePart(textprotoMap("text/html", "utf-8", ""))
	if err != nil {
		return fmt.Errorf("cannot create html part: %s", err)
	}
	if _, err = io.WriteString(part, message); err != nil {
		return fmt.Errorf("cannot write html part: %s", err)
	}

	for _, attachment := range attachments {
		f, err := os.Open(attachment)
		if err != nil {
			return fmt.Errorf("cannot open attachment %s: %s", attachment, err)
		}
		content, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			return fmt.Errorf("cannot read attachment %s: %s", attachment, err)
		}

		filename := filepath.Base(attachment)
		partHeader := textprotoMap("application/octet-stream", "", filename)
		partHeader.Set("Content-Transfer-Encoding", "base64")
		partHeader.Set("Content-Disposition",
			fmt.Sprintf("attachment;\r\n\tfilename=\"%s\"",
				mime.QEncoding.Encode("utf-8", filename)))

		part, err := writer.CreatePart(partHeader)
		if err != nil {
			return fmt.Errorf("cannot create attachment part for %s: %s", attachment, err)
		}
		b64 := base64.NewEncoder(base64.StdEncoding, part)
		if _, err = b64.Write(content); err != nil {
			return fmt.Errorf("cannot write attachment %s: %s", attachment, err)
		}
		b64.Close()
	}

	writer.Close()

	return smtp.SendMail(
		"smtp.gmail.com:587",
		authentication,
		sender,
		[]string{recipient},
		buf.Bytes(),
	)
}

func validateAttachments(attachments []string) error {
	for _, path := range attachments {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("attachment file not found: %s", path)
		}
	}
	return nil
}

func textprotoMap(contentType, charset, filename string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	ct := contentType
	if charset != "" {
		ct = fmt.Sprintf("%s; charset=%s", ct, charset)
	}
	h.Set("Content-Type", ct)
	if filename != "" {
		h.Set("Content-Disposition",
			fmt.Sprintf("inline;\r\n\tfilename=\"%s\"",
				mime.QEncoding.Encode("utf-8", filename)))
	}
	return h
}
