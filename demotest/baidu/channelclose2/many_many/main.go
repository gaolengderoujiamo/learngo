// 当有多个发送端和多个接收端时，无法再套用之前的方法。
// 因为关闭channel的操作只能执行一次，之前当接收端是一个的时候，是把接收端创造成一个发送端，
// 那么现在的情况，就可以考虑重新创造一个“发送端”，来做关闭通道的工作。
// 发送端和接收端都有终止channel通讯的条件，当随机值为0时发送端终止，当随机值为99999时接收端终止。
// 在没有达到条件的时候，moderator的goroutine中，toStop通道被阻塞，当有发送端或接收端向toStop中传值之后，
// stopCh通道的关闭接着执行，然后所有发送端和接收端都直接返回。同样，数据通道ch没有被显式关闭。

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MaxRandomNumber2 = 100000
	NumSenders       = 1000
	NumReceivers     = 10
)

var (
	wg2       sync.WaitGroup
	stoppedBy string
)

// 调度器
func moderator(stopCh chan struct{}, toStop chan string) {
	stoppedBy = <-toStop
	close(stopCh)
}

func senders(id int, stopCh chan struct{}, toStop chan string, ch chan int) {
	var value int
	for {
		value = rand.Intn(MaxRandomNumber2)
		if 0 == value {
			time.Sleep(2 * time.Second) // 这里睡眠两秒是为了让接收端也有一定概率先于发送端向toStop传值，从而更加公平一点
			select {
			case toStop <- fmt.Sprintf("sender#%d : %d\n", id, value):
			default:
			}
			return
		}

		select {
		case <-stopCh:
			return
		case ch <- value:
		default:
		}
	}
}

func receivers(id int, stopCh chan struct{}, toStop chan string, ch chan int) {
	defer wg2.Done()

	for {
		select {
		case <-stopCh:
			return

		case value := <-ch:
			if MaxRandomNumber2-1 == value {
				select {
				case toStop <- fmt.Sprintf("receiver#%d : %d\n", id, value):
				default:
				}
				return
			}
			fmt.Println(value)
		default:
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	ch := make(chan int, 100)
	stopCh := make(chan struct{})
	toStop := make(chan string, 1)

	go moderator(stopCh, toStop)

	for i := 0; i < NumSenders; i++ {
		go senders(i, stopCh, toStop, ch)
	}

	wg2.Add(NumReceivers)
	for i := 0; i < NumReceivers; i++ {
		go receivers(i, stopCh, toStop, ch)
	}
	wg2.Wait()

	fmt.Printf("Stopped by %s\n", stoppedBy)
}
