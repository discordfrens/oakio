package oakio

import (
	"log"

	"github.com/emersion/go-imap/backend/memory"
	"github.com/emersion/go-imap/server"
)

func NewServer() {
	be := memory.New()
	be.Login(nil, "hi@bun.red", "password")
	s := server.New(be)
	s.Addr = ":1143"
	s.AllowInsecureAuth = true
	log.Println("Starting IMAP server at localhost:1143")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}