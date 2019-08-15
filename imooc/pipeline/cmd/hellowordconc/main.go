package main

import (
	"fmt"
)

func main() {

	//var wg sync.WaitGroup
	c := make(chan string)
	for i := 0; i < 500; i++ {
		//wg.Add(1)
		go work(i, c)
		//wg.Done()
	}

	for {
		msg := <-c
		fmt.Println(msg)
	}
	//wg.Wait()

}

func work(i int, c chan string) {
	for {
		c <- fmt.Sprintf("Hello world from goroutine %d.\n", i)
	}
}
