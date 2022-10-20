package sendsms

import (
	"sync"
	"test/constants"
	"test/pkg/server/sendsms/model"
)

type SendMessage struct {
}

func NewSendMessage() *SendMessage {
	return &SendMessage{}
}

var jobs = make([]chan model.Task, constants.PartOfNum+1)

func (s *SendMessage) SendData(message model.Message, partOfLen int, restData int, wg *sync.WaitGroup) {
	numlist := *message.PhoneNum
	for i := 0; i < constants.PartOfNum; i++ {
		pointer := i
		jobs[pointer] = make(chan model.Task)
		Data := model.NewIntChan()
		Data = append(Data, numlist[i*partOfLen:(i+1)*partOfLen]...)
		go func() {
			for j := 0; j < partOfLen; j++ {
				jobs[pointer] <- model.Task{
					Args: Data[j],
					Do:   SendSms,
				}
			}
			close(jobs[pointer])
			wg.Done()
		}()
	}

	if restData != 0 {
		Data := model.NewIntChan()
		Data = append(Data, numlist[constants.PartOfNum*partOfLen:]...)
		jobs[constants.PartOfNum] = make(chan model.Task)
		go func() {
			for j := 0; j < len(Data); j++ {
				jobs[constants.PartOfNum] <- model.Task{
					Args: Data[j],
					Do:   SendSms,
				}
			}
			close(jobs[constants.PartOfNum])
			wg.Done()
		}()
	}
}

func (s *SendMessage) ReceiveData(message model.Message, restData int, wg *sync.WaitGroup) {
	for i := 0; i < constants.PartOfNum; i++ {
		pointer := i
		go func() {
			for {
				if v, ok := <-jobs[pointer]; !ok {
					break
				} else {
					_ = v.Do(v.Args, message.Content)
				}
			}
			wg.Done()
		}()
	}
	if restData != 0 {

		go func() {
			for {
				if v, ok := <-jobs[constants.PartOfNum]; !ok {
					break
				} else {
					_ = v.Do(v.Args, message.Content)
				}
			}
			wg.Done()
		}()
	}
}
