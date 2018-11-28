package main

import (
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
	statusParse    string        = " URL (statusParse) Некоректный адрес => "
	statusPing     string        = " PING (exec.command(ping)) Ошибка запроса => "
	succesPing     string        = " PING Успех => "
	respHeaderTime time.Duration = time.Second * 10 // Ожидание ответа http заголовка
)

// URLRequestGet - Запрос на указанный адрес.
// Возращает статус кода ответа (type string)
// В случаи ошибки статус ответа (type error)
func (u *URL) URLRequestGet() (string, error) {

	// Переменная client для сброса дефолтных значение конфигурации
	// клиента.
	// Задержка запроса Timeout (по умалчанию = ~ бесконечно)
	// приводит к зависание рутины.
	// Установить значение respHeaderTime в пределах от 5 - 10 секунд
	var client = &http.Client{
		Timeout: respHeaderTime,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: respHeaderTime,
			}).Dial,
			TLSHandshakeTimeout:   respHeaderTime,
			ResponseHeaderTimeout: respHeaderTime,
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
	return u.Name + "\n" + statusSucces + resp.Status, nil
}

// URLRequestPing - Ping указанного URL.
// Использует exec.Command возвращает статус (type string)
// В случаи ошибки статус ответа (type error)
func (u *URL) URLRequestPing() (string, error) {

	ur, err := url.Parse(u.URLandIP)
	if err != nil {
		return statusParse, err
	}
	_, err = exec.Command("ping", ur.Host, "-c 1", "-i 2", "-w 10").Output()
	if err != nil {
		return statusPing, err
	}
	//s := string(out[:])
	return succesPing + ur.Host, nil
}
