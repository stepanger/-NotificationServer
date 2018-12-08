package main

import (
	"fmt"
	"time"
)

// TimeStartTick - итерация запросов на сервер
// result принимает параметры в формате json
// second задержка цикла for
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
		effort := effort()

		if get {
			log, err := url.URLRequestGet()
			if err != nil {
				Error.Println(err)
				if effort == fail {
					NotifyLinux("Get Сервер не отвечает")
					gmail.SendingMessGmail("GET " + log + err.Error())
				}
			}

			fmt.Printf(log + "\n")
		}
		if ping {
			log, err := url.URLRequestPing()
			if err != nil {
				if effort == fail {
					NotifyLinux("Ping Сервер не отвечает")
					gmail.SendingMessGmail("PING " + log + err.Error())
				}
			}
			fmt.Printf(log + "\n")
		}
		if !get && !ping {
			fmt.Printf("\n\n ***Все запросы отключены!*** \n\n")
			return 0
		}
		time.Sleep(second)
	}
}
