package main

import (
	"time"
	. "readWeight/log1"
	"math/rand"
	"sync"
	"readWeight/cache"
)

type temp struct {
	index                  int
	list                   []int
	continueSameWeightFlag int
}

var (
	Flag = 0
	wait sync.WaitGroup
	t    = temp{0, make([]int, 100), 0}
)

func main() {
	wait.Add(1)

	go readSerial()
	Flag = 1

	wait.Wait()
}

func readSerial() {
	rand.Seed(time.Now().Unix())
	for {
		if Flag == 1 {
			var w = rand.Intn(2)
			Logger.Printf("获取体重:%dkg\n", w)
			isStability := continueSameWeight(w)
			if isStability {
				Logger.Printf("连续%d次 成功获取稳定体重%dkg", t.continueSameWeightFlag, w)
				t = temp{0, make([]int, 50), 0}
				cache.Weight = float32(w)
				Flag = 0

			}

		}

		time.Sleep(time.Second)
	}
}

func continueSameWeight(w int) bool {
	if t.index < len(t.list) {
		t.list[t.index] = w

		if t.index-1 < 0 {
			t.continueSameWeightFlag++
		} else if t.list[t.index-1] == t.list[t.index] {
			t.continueSameWeightFlag++
		} else {
			t.continueSameWeightFlag = 0
		}
		t.index++
		Logger.Printf("此次index[%d] 连续[%d]", t.index, t.continueSameWeightFlag)

	} else {
		t = temp{0, make([]int, 50), 0}
	}

	if t.continueSameWeightFlag >= 5 {
		return true
	}
	return false

}
