package main

import "os/exec"

// Failure - число неудачных попыток.
// С каждой итерацией увеличевает число до предела
// заданным effort, после сбрасывает x = 0
func Failure(effort int) func() int {
	var x int
	return func() int {
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
