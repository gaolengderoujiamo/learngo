package main

import (
	"fmt"
)

func fibo(c, stopch chan int) {
	x, y := 1, 1
	for {
		select {
		case <-stopch:
			return
		case c <- x:
			x, y = y, x+y
		}
	}
}

func main() {
	c := make(chan int)
	stopch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		close(stopch)
	}()

	fibo(c, stopch)

}
