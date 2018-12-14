package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Парсинг аргументов
	flag.Parse()

	// Лог приложения
	errorHandle, err := os.OpenFile("NS_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	infoHandle, err := os.OpenFile("NS_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer errorHandle.Close()
	defer infoHandle.Close()

	InitlogFile(errorHandle, infoHandle)
	//
	//
	//
	//
	// Конфигурация
	manifest, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(manifest), &result)
	//
	//
	//
	//
	//
	// Итерация запросов
	TimeStartTick(result)

}
