package main

import (
	"fmt"
)

type T int

func main() {
	done := make(chan T)
	close(done)
	fmt.Println(IsClosed(done))
}

func IsClosed(done chan T) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
