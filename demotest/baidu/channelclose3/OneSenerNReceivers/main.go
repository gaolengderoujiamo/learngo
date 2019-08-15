package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumReceivers = 100

	var wg sync.WaitGroup
	wg.Add(NumReceivers)

	dataCh := make(chan int, 100)

	// sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// the only sender can close the channel safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wg.Done()

			for v := range dataCh {
				log.Println(v)
			}
		}()
	}

	wg.Wait()
}
