package main

import (
	"fmt"
	"testing"
)

func TestNotifyLinux(t *testing.T) {
	NotifyLinux("test.go")
}

func TestFailure(t *testing.T) {

	effort := Failure(1000)

	for {
		effort := effort()
		if effort == 1000 {
			fmt.Println("1000")
			return
		}
	}
}

func TestURLRequestGet(t *testing.T) {
	url := &URL{
		"//google.com",
		"google inc",
	}

	log, err := url.URLRequestGet()
	fmt.Println("Запрос ", log, err)

}
