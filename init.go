package main

import "fmt"

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
		"209.185.108.134",
		"8000",
	}

	fmt.Println(url.URLRequestGet())
}
