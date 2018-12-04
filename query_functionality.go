package main

import "os/exec"

// Failure - число неудачных попыток.
// С каждой итерацией увеличевает число до предела
// заданным effort, после сбрасывает x = 0
func Failure(effort float64) func() float64 {
	var x float64
	return func() float64 {
		if effort == x {
			x = 0
		} else {
			x++
		}
		return x
	}
}

// NotifyLinux - Оповещание с помощью notify-send (Linux)
func NotifyLinux(status string) {
	header := "NotificationServer"
	cmd := exec.Command("notify-send", header, status)
	cmd.Run()
}
