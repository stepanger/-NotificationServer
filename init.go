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
		"192.168.1.1", //проверь с url
		"8000",
	}

	fmt.Println(url.URLRequestGet())
	fmt.Printf("%s", url.URLRequestPing())

	//out.Run() //.Output()

}
