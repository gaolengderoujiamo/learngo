package main

import (
	"fmt"
	"time"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occured:", err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is an error"))
	panic(fmt.Errorf("at: %s", time.Now()))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(123)
}

func main() {
	tryRecover()
}
