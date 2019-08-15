package channnelclose

import "sync"

type T int

type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (this *MyChannel) SafeClose() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if !this.closed {
		close(this.C)
		this.closed = true
	}
}

func (this *MyChannel) IsClosed() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.closed
}
