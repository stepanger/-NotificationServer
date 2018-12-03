package main

import (
	"os/exec"
)

func NotifyLinux(status string) {
	cmd := exec.Command("notify-send", status)
	cmd.Run()
}

func square() func() int {
	var x int = 0
	return func() int {
		x++
		return x
	}
}
