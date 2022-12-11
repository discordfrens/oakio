package smtp

import (
	"fmt"
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func Client() {
	fmt.Println("SMTP Client:")
	auth := sasl.NewPlainClient("", "hi@bun.red", "password")
	to := []string{
		"saigeeees@gmail.com",
	}
	msg := strings.NewReader(fmt.Sprintf("To: %v\r\n"+
		"Subject: discount Gophers!\r\n"+
		"\r\n"+
		"This is the email body.\r\n", to[0]))
	err := smtp.SendMail("192.168.0.9:143", auth, "hi@bun.red", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n")
}
