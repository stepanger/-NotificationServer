package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

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

	// перевод request_frequency в тип time.Duration, значение в секундах
	second := time.Duration(result["request_frequency"].(float64)) * time.Second

	timeStartTick(*url, second)
}

func timeStartTick(url URL, second time.Duration) {
	for {
		fmt.Println(url.URLRequestGet())
		time.Sleep(second)
	}
}
