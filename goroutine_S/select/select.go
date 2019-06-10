package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(id int, w worker) {
	for n := range w.in {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// worker done具体做的事情,由外边createWorker提供
		// 将sync.WaitGroup的Done()由worker的定义抽象出来
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) *
				time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	c1, c2 := generator(), generator()
	wg := sync.WaitGroup{}
	worker := createWorker(0, &wg)

	var values []int
	after := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker.in
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: // nil channel will yield
			values = values[1:]
			wg.Add(1)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("Queue len=", len(values))
		case <-after:
			fmt.Println("Bye")
			return
		}
	}
	wg.Wait()

}
