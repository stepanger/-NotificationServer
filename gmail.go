package main

import (
	"log"
	"net/smtp"
)

type gmail struct {
	sender     string //Отправитель
	passSender string //Пароль отправителя от почты gmail
	recipient  string //Получатель
}

const (
	subject  string = "Subject: NotificationServer\n\n"
	smtpServ string = "smtp.gmail.com"
	smtpPort string = ":587"
)

// SendingMessGmail - Отправка сообщения через почту Gmail
// Использует аторизацию с помощью PlainAuth и отправку через
// SendMail()
func (g *gmail) SendingMessGmail(mess string) {

	err := smtp.SendMail(smtpServ+smtpPort,
		smtp.PlainAuth("",
			g.sender,
			g.passSender,
			smtpServ),
		g.sender,
		[]string{g.recipient},
		[]byte(subject+mess))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
