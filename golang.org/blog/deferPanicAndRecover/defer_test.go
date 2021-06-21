package deferPanicAndRecover

import (
	"fmt"
	"testing"
)

// A deferred function's arguments are evaluated
//when the defer statement is evaluated.
func TestDefer_1(t *testing.T) {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// Deferred function calls are executed in
// Last In First Out order after the surrounding function returns.
func TestDefer_2(t *testing.T) {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

// Deferred functions may read and assign to the returning function's named return values.
func TestDefer_3(t *testing.T) {
	fmt.Println(c())
}
