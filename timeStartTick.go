package main

import (
	"fmt"
	"time"
)

func TimeStartTick(result map[string]interface{}, second time.Duration) int {

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
		if get {
			log, err := url.URLRequestGet()
			if err != nil {
				if effort() == fail {
					NotifyLinux("Сервер не отвечает")
					gmail.SendingMessGmail("Ошибка запроса сервер не отвечает <br>" + log)
				}
			}

			fmt.Printf(log)
		}
		if ping {
			fmt.Printf(url.URLRequestPing())
		}
		if !get && !ping {
			fmt.Printf("\n\n ***Все запросы отключены!*** \n\n")
			return 0
		}
		time.Sleep(second)
	}
}
