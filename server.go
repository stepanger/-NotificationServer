package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"time"
)

//URL - структура адреса
type URL struct {
	URLandIP string
	Name     string
}

const (
	statusError    string        = " HTTP Ошибка запроса => "
	statusSucces   string        = " HTTP Успех запроса => "
	respHeaderTime time.Duration = 10 // Ожидание ответа http заголовка
)

// URLRequestGet - Запрос на указанный адрес.
// Возращает статус кода ответа (type string)
// В случаи ошибки статус ответа (type error)
func (u *URL) URLRequestGet() (string, error) {

	// Нужен для сброса дефолтных значение клиента.
	// Задержка запроса Timeout (по умалчанию = 0)
	// тоесть бесконечно, установить в пределах от
	// 5 - 10 секунд
	var client = &http.Client{
		Timeout: respHeaderTime * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: respHeaderTime * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   respHeaderTime * time.Second,
			ResponseHeaderTimeout: respHeaderTime * time.Second,
		},
	}

	resp, err := client.Get(u.URLandIP)
	if err != nil {
		return u.Name + statusError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return u.Name + statusError, err
	}
	return u.Name + statusSucces + resp.Status, nil
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
	NotifyLinux("ping есть")
	return out
}
