package main

import (
	"os/exec"
)

// NotifyLinux - Оповещание с помощью notify-send (Linux)
func NotifyLinux(status string) {
	header := "NotificationServer"
	cmd := exec.Command("notify-send", header, status)
	cmd.Run()
}
