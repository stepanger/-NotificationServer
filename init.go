package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

/*
 Конфиг
 Адрес сервер или домен (множество)
 -- Ping по адресу
 -- Запрос Get, PUT,
 Частота проверки
 Уведомление на email
 Уведомление в вк
 Команда после сбоя + bash
*/
func main() {

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

	ticker := time.NewTicker(2000 * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			fmt.Println(url.URLRequestGet())
		}
	}()
	time.Sleep(time.Hour * 720)
	ticker.Stop()
}
