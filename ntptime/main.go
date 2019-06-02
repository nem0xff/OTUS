package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {

	fmt.Println(time.Now().Format("Системное время 15:04:05 Дата: 02.01.2006"))

	ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		panic("Что-то не так")
	}

	fmt.Println(ntptime.Format("NTP-сервер: 15:04:05 Дата: 02.01.2006"))
}
