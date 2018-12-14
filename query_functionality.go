package main

import (
	"io"
	"log"
	"os/exec"
)

// InitlogFile => log.New()
var (
	Error *log.Logger
	Info  *log.Logger
)

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

// InitlogFile - лог выполнения программы
// NS ERROR: _ гггг/мм/дд 01:59:59 error
// NS INFO: _ гггг/мм/дд 00:00:00 string
func InitlogFile(errorHandle, infoHandle io.Writer) {
	Error = log.New(errorHandle, "NS ERROR: _ ", log.LstdFlags)
	Info = log.New(infoHandle, "NS INFO: _ ", log.LstdFlags)
}
