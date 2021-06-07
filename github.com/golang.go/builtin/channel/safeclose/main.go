package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type T struct{}

func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *MyChannel) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}

func main() {
	randomSelect()
}

func randomSelect() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val:%d\n", val)
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case ch <- i:
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		}

		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}

func closeChan() {
	c := make(chan int, 1)

	go func() {
		c <- 1
		log.Println("send first")
		c <- 2
		log.Println("send second")
		close(c)
		return
	}()

	for data := range c {
		time.Sleep(5 * time.Second)
		fmt.Println(data)
	}
}

func channelCloseSelect() {
	c := make(chan int)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			select {
			//case <-time.Tick(time.Second):
			//	c <- rand.Intn(100)
			case <-time.After(3 * time.Second):
				close(done)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case data := <-c:
				fmt.Println(data)
			case <-done:
				log.Println("done")
				return
			}
		}
	}()

	wg.Wait()
}
