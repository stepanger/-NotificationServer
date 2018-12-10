package main

import (
	"net/smtp"
)

// Gmail Получатель почты
type Gmail struct {
	MailRecipient  string //Получатель
	Mailer         string //Отправитель .bashrc
	SenderPassword string //Пароль отправителя .bashrc
}

const (
	subject    string = "Subject: NotificationServer\r\n\r\n"
	smtpServer string = "smtp.gmail.com"
	smtpPort   string = ":587"
)

// SendingMessGmail - Отправка сообщения через почту Gmail
// Использует аторизацию с помощью PlainAuth и отправку через
// SendMail().
func (g *Gmail) SendingMessGmail(mess string, authGmail smtp.Auth) {

	err := smtp.SendMail(smtpServer+smtpPort,
		authGmail,
		g.Mailer,
		[]string{g.MailRecipient},
		[]byte("To: Admin\r\n"+subject+mess))

	if err != nil {
		Error.Println("smtp error: ", err)
		return
	}
}

// AuthenticationUser - Авторизация пользователя Gmail
// Авторизация береться из .bashrc
func (g *Gmail) AuthenticationUser() smtp.Auth {
	authGmail := smtp.PlainAuth(
		"",
		g.Mailer,
		g.SenderPassword,
		smtpServer,
	)
	return authGmail
}
