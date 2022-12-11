package main

import (
	"oakio-mail-client/main/imap"
	"oakio-mail-client/main/smtp"
)

func main() {
	smtp.Client()
	imap.Client()
}
