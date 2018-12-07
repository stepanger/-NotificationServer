package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	// Лог приложения
	fileNS, err := os.OpenFile("logNS.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer fileNS.Close()
	logNS := log.New(fileNS, "NS _ ", log.LstdFlags)

	// Конфигурация
	content, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		logNS.Println(err)
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)

	// перевод request_frequency в тип time.Duration, значение в секундах
	second := time.Duration(result["request_frequency"].(float64)) * time.Second

	// Итерация запросов
	TimeStartTick(result, second)

}
