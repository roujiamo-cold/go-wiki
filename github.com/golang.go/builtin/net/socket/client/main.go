package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			dial()
		}
		fmt.Printf("happen")
	}
}

func dial() {
	conn, err := net.Dial("tcp", ":6010")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "hello server\n")

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", data)
}
