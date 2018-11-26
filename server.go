package main

import (
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

type URL struct {
	URLandIP, Name string
}

// URLRequestGet - Запрос на указанный адрес.
// Возращает статус кода ответа (type string)
func (u *URL) URLRequestGet() (string, bool) {
	response, err := http.Get(u.URLandIP)
	if err != nil {
		return u.Name + " HTTP Ошибка запроса => URLRequestGet()", true
	}
	defer response.Body.Close()
	return u.Name + " HTTP Запрос => " + response.Status, false
}

// URLRequestPing - Ping указанного URL.
// Использует exec.Command возвращает статус (type []byte)
func (u *URL) URLRequestPing() []byte {

	ur, err := url.Parse(u.URLandIP)
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("ping", ur.Host, "-c 1", "-i 2", "-w 10").Output()
	if err != nil {
		log.Fatalf("STATUS ERROR %s", err)
	}
	return out
}
