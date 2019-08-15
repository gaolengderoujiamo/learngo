package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sender: 1, receivers: M
const (
	MaxRandomNumber = 100000
	NumReceivers1   = 100
)

var wg sync.WaitGroup

func sender(ch chan int) {
	var value int
	for {
		value = rand.Intn(MaxRandomNumber)
		if 0 == value { // 当随机生成的数为0时，关闭通道
			fmt.Printf("Closed for number: %d\n", value)
			close(ch)
			return
		} else {
			ch <- value
		}
	}
}

func receiver(ch chan int) {
	defer wg.Done()

	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	ch := make(chan int, 100)

	go sender(ch)

	wg.Add(NumReceivers1)
	for i := 0; i < NumReceivers1; i++ {
		go receiver(ch)
	}
	wg.Wait()
}
