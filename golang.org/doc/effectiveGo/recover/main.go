package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type myErr struct {
	msg string
	err error
}

func (m *myErr) Error() string {
	return fmt.Sprintf("myErr: msg: %s, err: %v", m.msg, m.err)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//panic("this is my panic")

	err := errors.New("new error")

	myErr := err.(*myErr)

	fmt.Printf("%v", myErr)
}

// -------
type work interface{}

func do(w *work) {
	panic("not implements")
}

func server(workChan <-chan *work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

// -------
func do1() {
	//regexp.Compile()
}
