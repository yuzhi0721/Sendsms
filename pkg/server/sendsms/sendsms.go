package sendsms

import (
	"fmt"
	"sendsms/pkg/sendsms/pack"
	"sendsms/pkg/sendsms/sendsms"
	"sendsms/pkg/server/sendsms/utils"
	"sync"
)

var wg sync.WaitGroup

func init() {

}
func SendSms() {
	sendMessage := sendsms.NewSendMessage()
	content := pack.DataSrcContent()
	numlist := pack.DataSrcNum()
	Message := pack.NewMessage(content, numlist)
	x, y := utils.NumOfGo(*numlist)
	utils.WgNum(y, &wg)
	sendMessage.SendData(*Message, x, y, &wg)
	sendMessage.ReceiveData(*Message, y, &wg)
	fmt.Println("rebase")
	wg.Wait()
}
