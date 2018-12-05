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
	statusError    string        = " HTTP (URLRequestGet) Ошибка запроса => "
	statusSucces   string        = " HTTP Успех => "
	statusParse    string        = " URL (statusParse) Некоректный адрес => "
	statusPing     string        = " PING (exec.command(ping)) Ошибка запроса => "
	succesPing     string        = " PING Успех => "
	respHeaderTime time.Duration = time.Second * 10 // Ожидание ответа http заголовка
)

// URLRequestGet - Выполняет GET запрос на указанный адрес.
// Возращает статус кода ответа (type string)
// В случаи ошибки статус ответа + (error)
func (u *URL) URLRequestGet() (string, error) {

	// Проверка адреса
	_, err := url.Parse(u.URLandIP)
	if err != nil {
		return statusParse + " " + u.URLandIP, err
	}

	// Переменная client для сброса дефолтных значение
	// конфигурации клиента. Задержка запроса
	// Timeout (по умалчанию = ~ бесконечно)
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

	// Запрос на сервер
	resp, err := client.Get(u.URLandIP)
	if err != nil {
		return u.Name + statusError, err
	}

	defer resp.Body.Close()

	// Сравнение с статус кода 200
	if resp.StatusCode != http.StatusOK {
		return u.Name + statusError, err
	}
	return u.Name + statusSucces + resp.Status, nil
}

// URLRequestPing - ping указанного URL.
// Использует exec.Command возвращает статус (type string)
// В случаи ошибки статус ответа (type error)
func (u *URL) URLRequestPing() (string, error) {

	// Парсинг для коректности url и сохранения домена
	urlParse, err := url.Parse(u.URLandIP)
	if err != nil {
		return statusParse + " " + u.URLandIP, err
	}

	_, err = exec.Command("ping", urlParse.Host, "-c 1", "-i 2", "-w 10").Output()
	if err != nil {
		return u.Name + statusPing, err
	}
	return u.Name + succesPing + "OK", nil
}
