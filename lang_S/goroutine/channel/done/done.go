package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

// done为function, 由createWorker进行配置
// doWorker仅负责调用worker的done()方法
// 做到doWorker与内部done()所做的事情抽象掉
type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
