package sendsms

import (
	"sync"
	"test/pkg/sendsms/pack"
	"test/pkg/sendsms/sendsms"
	"test/pkg/server/sendsms/utils"
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
	wg.Wait()
}
