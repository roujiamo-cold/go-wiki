package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestFmtErrorf(t *testing.T) {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

	errXiaoming := errors.New("xiaoming error")
	errWrap := fmt.Errorf("wrap error1: %w", errXiaoming)
	fmt.Println(errors.Is(errWrap, errXiaoming))
	fmt.Println(errors.Unwrap(errWrap))

}
