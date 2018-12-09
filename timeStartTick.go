package main

import (
	"time"
)

// TimeStartTick - итерация запросов на сервер
// result принимает параметры в формате json
// second задержка цикла for
func TimeStartTick(result map[string]interface{}) int {

	second := time.Duration(result["request_frequency"].(float64)) * time.Second

	url := &URL{
		result["http_url_host"].(string),
		result["name_host"].(string),
	}

	gmail := &Gmail{
		result["gmail_notification"].(string),
	}

	ping := result["reguest_ping"].(bool)
	get := result["reguest_http_get"].(bool)
	fail := result["failed_attempts"].(float64)

	effort := Failure(fail)

	for {
		effort := effort()

		if get {
			log, err := url.URLRequestGet()
			if err != nil {
				Error.Println(log, err)
				if effort == fail {
					NotifyLinux("Get Сервер не отвечает")
					gmail.SendingMessGmail("GET " + log + err.Error())
				}
			} else {
				Info.Println(log)
			}
		}
		if ping {
			log, err := url.URLRequestPing()
			if err != nil {
				Error.Println(log, err)
				if effort == fail {
					NotifyLinux("Ping Сервер не отвечает")
					gmail.SendingMessGmail("PING " + log + err.Error())
				}
			} else {
				Info.Println(log)
			}
		}
		if !get && !ping {
			Info.Println("***Все запросы отключены!***")
			return 0
		}
		time.Sleep(second)
	}
}
