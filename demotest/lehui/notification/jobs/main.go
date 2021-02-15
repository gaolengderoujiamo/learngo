package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething(id int, c chan int) {
	defer wg.Done()
	i := 0
	for {
		time.Sleep(time.Duration(rand.Intn(1500)) *
			time.Millisecond)
		i = i + 1
		fmt.Printf("%d - %d\n", id, i)

		select {
		case <-c:
			return
		default:

		}
	}
}

func shutdown(c chan int) {
	fmt.Println("shotduwn all the goroutine.")
	close(c)
}

func main() {
	wg.Add(2)
	c := make(chan int)
	for i := 0; i < 2; i++ {
		go doSomething(i, c)
	}

	time.Sleep(5 * time.Second)
	shutdown(c)
	wg.Wait()
}
