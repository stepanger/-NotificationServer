package main

import (
	"encoding/json"
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

	/////////////////////////////////////////////////////////////////////////

	// перевод request_frequency в тип time.Duration, значение в секундах
	second := time.Duration(result["request_frequency"].(float64)) * time.Second

	/////////////////////////////////////////////////////////////////////////////

	//	timeStartTick(*url, *gmail, second, getIP)
	TimeStartTick(result, second)

}
