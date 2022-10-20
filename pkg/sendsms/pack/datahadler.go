package pack

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"test/constants"
	"test/pkg/server/sendsms/model"
	"time"
)

func Content() string {
	contentpath := fmt.Sprintf("%s%s/content", constants.FilePath, time.Now().Format("2006010203"))
	data, err := ioutil.ReadFile(contentpath)
	if err != nil {
		return ""
	}
	return string(data)
}

func NumList() *[]int {
	var numslice []int
	numberpath := fmt.Sprintf("%s%s/number", constants.FilePath, time.Now().Format("2006010203"))
	numberfile, err := os.Open(numberpath)
	if err != nil {
		return nil
	}
	defer numberfile.Close()
	br := bufio.NewReader(numberfile)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		num, err := strconv.Atoi(string(line))
		if err != nil {
			return nil
		}
		numslice = append(numslice, num)
	}
	return &numslice
}

func DataSrcContent() string {
	return Content()
}

func DataSrcNum() *[]int {
	return NumList()
}

func NewMessage(c string, numList *[]int) *model.Message {
	return &model.Message{
		PhoneNum: numList,
		Content:  c,
	}
}
