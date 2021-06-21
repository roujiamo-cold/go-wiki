package testingTechniques

import (
	"fmt"
	"os"
)

func Crasher() {
	fmt.Println("Going down in flames!")
	os.Exit(1)
}
