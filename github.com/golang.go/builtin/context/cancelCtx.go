package context

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

type Value int

func DoSomething(ctx context.Context) (Value, error) {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(100)
	if intn == 99 {
		return 0, errors.New("99 error")
	}

	time.Sleep(500 * time.Millisecond)
	return Value(intn), nil
}

func cancelStream(ctx context.Context, out chan<- Value) error {
	for {
		v, err := DoSomething(ctx)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- v:
		}
	}
}

func CancelCtxTest() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	ch := make(chan Value)

	go func(fn context.CancelFunc) {
		// 5s 后调用cancel()
		time.Sleep(5 * time.Second)
		cancelFunc()
	}(cancelFunc)

	go func() {
		if err := cancelStream(ctx, ch); err != nil {
			fmt.Printf("stream error = %v\n", err)
			close(ch)
			return
		}
	}()

	for d := range ch {
		fmt.Println(d)
	}

}
