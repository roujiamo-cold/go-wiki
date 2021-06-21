package deferPanicAndRecover

import (
	"fmt"
	"testing"
)

func TestRecover_3(t *testing.T) {
	f()
	fmt.Println("Returned normally from f.")
}
