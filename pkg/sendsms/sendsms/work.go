package sendsms

import "fmt"

func SendSms(mobile int, content string) error {
	fmt.Printf("send %s to %v\n", content, mobile)
	return nil
}
