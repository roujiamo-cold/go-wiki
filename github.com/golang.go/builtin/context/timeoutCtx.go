package context

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

var timeTemplate = "2006-01-02 15:04:05"

func generator() (int, error) {
	time.Sleep(500 * time.Millisecond)
	intn := rand.Intn(100)
	if intn == 99 {
		return 0, errors.New("99 error")
	}
	return intn, nil
}

func timeoutStream(ctx context.Context, out chan<- int) error {
	for {
		i, err := generator()
		if err != nil {
			return errors.Wrap(err, "timeoutStream: generator error")
		}

		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "timeoutStream: ctx.Done()")
		case out <- i:
			deadline, ok := ctx.Deadline()
			log.Printf("ctx deadline %v, %v", deadline.Format(timeTemplate), ok)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TimeoutCtxTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		if err := timeoutStream(ctx, ch); err != nil {
			close(ch)
			fmt.Printf("%v\n", err)
			return
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
