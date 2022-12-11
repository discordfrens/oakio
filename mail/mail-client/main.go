package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {
	auth := sasl.NewPlainClient("", "username", "password")
	to := []string{
		"saigeeees@gmail.com",
	}
	msg := strings.NewReader(fmt.Sprintf("To: %v\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n", to[0]))
	err := smtp.SendMail("localhost:143", auth, "test@saige.systems", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
