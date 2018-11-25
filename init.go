package main

import (
	"fmt"
	"os/exec"
)

// Адрес сервер или домен (множество)
// -- Ping по адресу
// -- Запрос Get, PUT,
// Частота проверки
// Уведомление на email
// Уведомление в вк
// Команда после сбоя + bash

func main() {

	url := &URL{
		"https://www.google.ru",
		"google",
	}

	fmt.Println(url.URLRequestGet())
	fmt.Printf("%s", url.URLRequestPing())

	cmd := exec.Command("notify-send", "string")
	cmd.Run()

}
