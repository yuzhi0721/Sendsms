package utils

import (
	"sync"
	"test/constants"
)

func NumOfGo(x []int) (int, int) {
	return len(x) / constants.PartOfNum, len(x) % constants.PartOfNum
}

func WgNum(x int, wg *sync.WaitGroup) {
	if x != 0 {
		wg.Add(2*constants.PartOfNum + 2)
	} else {
		wg.Add(2 * constants.PartOfNum)
	}
}
