package oakio

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	// "os"
	"time"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username != "hi@bun.red" || password != "password" {
		return nil, errors.New("Invalid username or password")
	}
	return &Session{}, nil
}
// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

// A Session is returned after successful login.
type Session struct{}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
		err = smtp.SendMail("smtp.gmail.com:587", nil, "hi@bun.red", []string{"saigeeees@gmail.com"}, r)
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func NewServer() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":143"
	s.Domain = "localhost"
	s.ReadTimeout = 40 * time.Second
	s.WriteTimeout = 40 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true
	// s.Debug = os.Stdout

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}