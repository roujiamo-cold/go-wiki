package main

import (
	"context"
	"fmt"

	myctx "github.com/roujiamo-cold/go-wiki/github.com/golang.go/builtin/context"
)

var ctx context.Context

func main() {
	myctx.CancelCtxTest()
	//myctx.TimeoutCtxTest()
}

// userTest 1
func userTest() {
	var in myctx.User

	ctx := context.Background()
	u := &myctx.User{
		Name: "xiaoming",
	}
	newContext := myctx.NewContext(ctx, u)

	fmt.Println(myctx.FromContext(newContext))

	fmt.Println(ctx)
	fmt.Println(in.Name)
}
