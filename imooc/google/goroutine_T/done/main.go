package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(i, w)
	return w
}

func doWork(i int, w worker) {
	for v := range w.in {
		fmt.Printf("%c\n", v)
		w.done()
	}
}

func main() {
	var wg sync.WaitGroup

	workers := []worker{}
	for i := 0; i < 10; i++ {
		workers = append(workers, createWorker(i, &wg))
	}

	// group 1
	for i := 0; i < 10; i++ {
		wg.Add(1)
		workers[i].in <- 'a' + i
	}

	// group 2
	for i := 0; i < 10; i++ {
		wg.Add(1)
		workers[i].in <- 'A' + i
	}

	wg.Wait()

}
