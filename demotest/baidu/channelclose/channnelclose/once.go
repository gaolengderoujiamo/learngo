package channnelclose

import "sync"

type MyChannelOnce struct {
	C    chan T
	once sync.Once
}

func NewMyChannelOnce() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (this *MyChannelOnce) SafeClose() {
	this.once.Do(func() {
		close(this.C)
	})
}
