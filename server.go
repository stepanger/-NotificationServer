package main

import (
	"log"
	"net/http"
	"os/exec"
)

// URL - структура адреса
type URL struct {
	Name, IP, Port string
}

// URLRequestGet - Запрос на указанный адрес.
// Возращает статус кода ответа (type string)
func (u *URL) URLRequestGet() string {
	response, err := http.Get(u.Name)
	if err != nil {
		return "Error get !"
	}

	defer response.Body.Close()
	return response.Status

}

// URLRequestPing - Ping указанного IP.
// Использует exec.Command возвращает статус (type []byte)
func (u *URL) URLRequestPing() []byte {
	out, err := exec.Command("ping", u.IP, "-c 3", "-i 2", "-w 10").Output()

	if err != nil {
		log.Fatalf("STATUS ERROR %s", err)
	}

	return out

}

func (u *URL) URLRequestPUT() string { return "" }
