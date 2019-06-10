package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() <-chan int {
	c := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func worker(id int, c chan int) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, v)
	}
}

func createWorker(id int) chan<- int {
	w := make(chan int)
	go worker(id, w)
	return w
}

func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)

	var values []int
	after := time.After(10 * time.Second)
	tick := time.Tick(time.Second) // 定时任务
	for {
		// nil channel的特性,nil channel的case会阻塞
		var activeChan chan<- int
		var activeVal int
		if len(values) > 0 {
			activeChan = worker
			activeVal = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeChan <- activeVal:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 800ms没有收到数据,timeout
			fmt.Println("Time out!")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-after: // 程序10s自动退出
			fmt.Println("Bye!")
			return
		}
	}
}
