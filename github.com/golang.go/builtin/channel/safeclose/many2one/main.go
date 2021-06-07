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

	// ...
	const MaxRandomNumber = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its receivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(i int) {
			for {
				// The first select here is to try to exit the goroutine
				// as early as possible. In fact, it is not essential
				// for this example, so it can be omitted.
				select {
				case <-stopCh:
					log.Printf("[%d] first select receive stopCh\n", i)
					return
				default:
				}

				// Even if stopCh is closed, the first branch in the
				// second select may be still not selected for some
				// loops if the send to dataCh is also unblocked.
				// But this is acceptable, so the first select
				// can be omitted.
				select {
				case <-stopCh:
					log.Printf("[%d] second select receive stopCh\n", i)
					return
				case dataCh <- rand.Intn(MaxRandomNumber):

				}
			}
		}(i)
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				// The receiver of the dataCh channel is
				// also the sender of the stopCh cahnnel.
				// It is safe to close the stop channel here.
				close(stopCh)
				return
			}

			//log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()

	time.Sleep(5 * time.Second)
}
