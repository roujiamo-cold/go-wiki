package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func main() {
	ch := make(chan error)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- errors.New("done")
	}()

	fmt.Println(<-ch)
}
