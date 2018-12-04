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
$ echo $GmailUser
```
В терминале должно отображаться установленное значение
