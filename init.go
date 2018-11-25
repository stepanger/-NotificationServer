package main

import (
	"fmt"
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

	NotifyLinux("Ошибка подключения сервер не отвечает")

}
