package main

import (
	"context"
	"fmt"

	"github.com/roujiamo-cold/go-wiki/github.com/golang.go/builtin/context/user"
)

var ctx context.Context

func main() {

	var in user.User

	ctx := context.Background()
	u := &user.User{
		Name: "xiaoming",
	}
	newContext := user.NewContext(ctx, u)

	fmt.Println(user.FromContext(newContext))

	fmt.Println(ctx)
	fmt.Println(in.Name)
}
