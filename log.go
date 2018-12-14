package main

import (
	"io"
	"log"
)

// InitlogFile - лог выполнения программы
// NS ERROR: _ гггг/мм/дд 01:59:59 error
// NS INFO: _ гггг/мм/дд 00:00:00 string
func InitlogFile(errorHandle, infoHandle io.Writer) {
	Error = log.New(errorHandle, "NS ERROR: _ ", log.LstdFlags)
	Info = log.New(infoHandle, "NS INFO: _ ", log.LstdFlags)
}
