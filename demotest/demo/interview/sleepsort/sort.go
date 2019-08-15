package sleepsort

import (
	"fmt"
	"time"
)

func doWork(i int, c chan int) {
	go func() {
		time.Sleep(10000 * time.Millisecond)
		fmt.Println(i)
		c <- i
	}()
}

func SleepSort(data []int, n int) {
	c := make(chan int, len(data))
	for _, v := range data {
		doWork(v, c)
	}

	for i := 0; i < n; i++ {
		<-c
	}
}
