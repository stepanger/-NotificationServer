package main

import (
	"net/smtp"
	"os"
)

// Gmail Получатель почты
type Gmail struct {
	recipient string //Получатель
}

const (
	subject  string = "Subject: NotificationServer\n\n"
	smtpServ string = "smtp.gmail.com"
	smtpPort string = ":587"
)

// SendingMessGmail - Отправка сообщения через почту Gmail
// Использует аторизацию с помощью PlainAuth и отправку через
// SendMail(). Значение отправителя береться из .bashrc
func (g *Gmail) SendingMessGmail(mess string) {

	err := smtp.SendMail(smtpServ+smtpPort,
		smtp.PlainAuth("",
			os.Getenv("GmailUser"),
			os.Getenv("GmailPass"),
			smtpServ),
		os.Getenv("GmailUser"),
		[]string{g.recipient},
		[]byte(subject+mess))

	if err != nil {
		Error.Println("smtp error: ", err)
		return
	}
}
