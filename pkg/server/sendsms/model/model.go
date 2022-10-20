package model

type Task struct {
	Args int
	Do   func(int, string) error
}

type Message struct {
	PhoneNum *[]int
	Content  string
}

var Jobs []chan Task

func NewIntChan() []int {
	return make([]int, 0)
}
