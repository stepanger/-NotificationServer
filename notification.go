package main

import (
	"os/exec"
)

func NotifyLinux(status string) {
	cmd := exec.Command("notify-send", status)
	cmd.Run()
}
