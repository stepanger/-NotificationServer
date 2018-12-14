package main

import "flag"

var (
	logCommand  = flag.Bool("L", false, "Вывод логов в командной строке: по умолчанию false")
	httpRequest = flag.String("U", "", "HTTP запрос на указанный URL")
	pingRequest = flag.String("P", "", "Ping запрос на указанный URL")
)
