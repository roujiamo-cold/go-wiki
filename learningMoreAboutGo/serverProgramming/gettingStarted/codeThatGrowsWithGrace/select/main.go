package main

import (
	"fmt"
	"time"
)

func main() {
	//boom := time.After(time.Second * 1)
	ticker_1s := time.NewTicker(time.Second)
	ticker_250ms := time.NewTicker(time.Millisecond * 250)
	for {
		select {
		case <-ticker_1s.C:
			fmt.Println("boom!")
		case <-ticker_250ms.C:
			fmt.Println("tick")
		}
	}
}
