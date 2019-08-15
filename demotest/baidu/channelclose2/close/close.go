package close

import "sync"

type MyChannel struct {
	C    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan int)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

// ------------------------------------------

type MyChannel1 struct {
	C      chan int
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel1() *MyChannel {
	return &MyChannel{C: make(chan int)}
}

func (mc *MyChannel1) SafeClose() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *MyChannel1) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
