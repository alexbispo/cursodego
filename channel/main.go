package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go pong(canal)
	go ping(canal)
	go juiz(canal)

	var entrada string
	fmt.Scanln(&entrada)
}

func juiz(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func ping(canal chan string) {
	for {
		canal <- "Pong"
	}
}

func pong(canal chan string) {
	for {
		canal <- "Ping"
	}
}
