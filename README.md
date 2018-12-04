# NotificationServer
```sh
Linux 4.15.0-39-generic #42-Ubuntu SMP Tue  x86_64
```

#### Стадия (Pre-Alpha)
[![Go Report Card](https://goreportcard.com/badge/github.com/stepanger/NotificationServer)](https://goreportcard.com/report/github.com/stepanger/NotificationServer)

Проверяет есть ли отклик от сервера или web-приложения. Manifest.json основной конфик приложения.
- Запрос GET
- Ping адреса
- Уведомление при сбое Gmail, Notify-send

#### Установка (Linux)
###### Переменные окружения для Gmail
Установить переменную окружения в файл .bashrc домашней дериктории
в конец файла
```sh
#.bashrc
export GmailUser="...@gmail.com"
export GmailPass="password"
```
Перезапустите терминал и введите команду.

```sh
@ echo $GmailUser
```
В терминале должно отображаться установленное значение

###### Параметры manifest.json

| Параметр         |Тип       |Описание |
|------------------|----------|-------- |
|name_host         | `string` |Наименование сервера |
|http_url_host     | `string` |Адресная строка URL |
|reguest_ping      | `bool`   |Если true то выполнит ping запрос, false отменяет запрос  |
|reguest_http_get  | `bool`   |Если true то выполнит get запрос |
|request_frequency | `float64`|С какой периодичностью выполнять запрос в секундах |
|failed_attempts   | `float64`|Число попыток, после чего уведомление по почте |
|gmail_notification| `string` |Email адресс получателя  |

Пример:
```json
{
  "name_host": "google inc",
  "http_url_host": "https://www.google.com/",
  "reguest_ping": false,
  "reguest_http_get": true,
  "request_frequency": 1,
  "failed_attempts": 10,
  "gmail_notification": "creadtive472@gmail.com"
}
```
