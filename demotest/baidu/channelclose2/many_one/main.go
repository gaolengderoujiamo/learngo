///当发送端为多个的时候，为了避免多次关闭channel，可以考虑增加一个作为标志的channel，
// 当需要关闭channel的时候，通过关闭标志channel，通知多个发送端结束工作从而停止了发送工作。
// 换一个角度去想，就标志channel而言，唯一的接收端是它的实际发送者，因此依然遵循通道关闭原则。
// 所以实际上，传输数据的channel并没有显式关闭。
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MaxRandomNumber1 = 100000
	NumSender1       = 100
)

var wg1 sync.WaitGroup

func sender1(ch chan int, stopCh chan struct{}) {
	var value int
	for {
		value = rand.Intn(MaxRandomNumber1)

		select {
		case <-stopCh:
			return
		case ch <- value:

		}
	}
}

func receiver1(ch chan int, stopCh chan struct{}) {
	defer wg1.Done()
	for value := range ch {
		if MaxRandomNumber1-1 == value {
			close(stopCh)
			fmt.Printf("From receiver1 closed by number: %d\n", value)
			return
		}

		fmt.Println(value)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	ch := make(chan int, 100)
	stopCh := make(chan struct{})

	for i := 0; i < NumSender1; i++ {
		go sender1(ch, stopCh)
	}

	wg1.Add(1)
	go receiver1(ch, stopCh)
	wg1.Wait()
}
