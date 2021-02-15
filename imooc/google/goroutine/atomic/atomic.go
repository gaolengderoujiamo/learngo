package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increament() {
	fmt.Println("Safe increament")
	func() { // 原子性的代码块
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increament()
	go func() {
		a.increament()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
