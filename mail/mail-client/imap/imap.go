package imap

import (
	"fmt"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func Client() {
	fmt.Println("IMAP Client:")
	log.Println("Connecting to imap server....")
	c, err := client.Dial("192.168.0.9:1143")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()
	if err := c.Login("hi@bun.red", "password"); err != nil {
		log.Fatal(err)
	}

	log.Println("Logged in as hi@bun.red")
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n")
}
