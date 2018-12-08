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
	errorHandle, err := os.OpenFile("NS_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer errorHandle.Close()

	infoHandle, err := os.OpenFile("NS_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer infoHandle.Close()

	InitlogFile(errorHandle, infoHandle)
	//
	//
	//
	//
	// Конфигурация
	content, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)
	//
	//
	//
	//
	// перевод request_frequency в тип time.Duration, значение в секундах
	second := time.Duration(result["request_frequency"].(float64)) * time.Second
	//
	//
	//
	//
	//
	// Итерация запросов
	TimeStartTick(result, second)

}
