package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Конфиг
// Адрес сервер или домен (множество)
// -- Ping по адресу
// -- Запрос Get, PUT,
// Частота проверки
// Уведомление на email
// Уведомление в вк
// Команда после сбоя + bash

func main() {

	//fmt.Println(url.URLRequestGet())
	//fmt.Printf("%s", url.URLRequestPing())

	content, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)

	url := &URL{
		result["http_url_host"].(string),
		result["name_host"].(string),
	}

	fmt.Println(url.URLRequestGet())
	fmt.Printf("%s", url.URLRequestPing())

}
