package main

import (
	"net/http"
	"os"
)

// URL - структура адреса
type URL struct {
	Name string
	IP   string
	Port string
}

// URLRequestGet - Запрос на указанный адрес.
// Возращает статус кода ответа (type string)
func (u *URL) URLRequestGet() string {
	response, err := http.Get(u.Name)
	if err != nil {
		os.Exit(1)
		return "Error get !"
	}

	defer response.Body.Close()
	return response.Status

}

func (u *URL) URLRequestPing() string { return "" }
func (u *URL) URLRequestPUT() string  { return "" }
