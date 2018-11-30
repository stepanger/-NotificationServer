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

	getIP := map[string]bool{
		"Ping": result["reguest_ping"].(bool),
		"Get":  result["reguest_http_get"].(bool),
	}

	timeStartTick(*url, second, getIP)

}

func timeStartTick(url URL, second time.Duration, getIP map[string]bool) int {

	for {
		switch {
		case getIP["Ping"] && getIP["Get"]:
			fmt.Printf(url.URLRequestGet())
			fmt.Printf(url.URLRequestPing())
		case getIP["Ping"]:
			fmt.Println(url.URLRequestPing())
		case getIP["Get"]:
			fmt.Println(url.URLRequestGet())
		default:
			fmt.Printf("\n\n ***Нечего отправлять!*** \n\n")
			return 0
		}
		time.Sleep(second)
	}
}
