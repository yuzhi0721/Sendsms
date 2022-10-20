package main

import (
	"fmt"
	"sendsms/pkg/server/sendsms"
	"time"
)

func main() {
	sendsms.SendSms()
	fmt.Println(time.Now().Format("2006010203"))
}
