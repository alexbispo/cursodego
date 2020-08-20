package main

import (
	"fmt"
	"time"
)

type mensagemPing struct {
	body string
}

type mensagemPong struct {
	body string
}

func main() {
	canalPing := make(chan mensagemPing)
	canalPong := make(chan mensagemPong)

	go func() { canalPing <- mensagemPing{body: "Start"} }()
	go ping(canalPing, canalPong)
	go pong(canalPong, canalPing)

	var entrada string
	fmt.Scanln(&entrada)
}

func ping(canalPing chan<- mensagemPing, canalPong <-chan mensagemPong) {
	for {
		select {
		case msg := <-canalPong:
			fmt.Println(msg.body)
			canalPing <- mensagemPing{body: "Ping"}
		default:
			fmt.Println("Mensagem não recebida")
		}

		time.Sleep(time.Second)
	}
}

func pong(canalPong chan<- mensagemPong, canalPing <-chan mensagemPing) {
	for {
		select {
		case msg := <-canalPing:
			fmt.Println(msg.body)
			canalPong <- mensagemPong{body: "Pong"}
		default:
			fmt.Println("Mensagem não recebida")
		}

		time.Sleep(time.Second)
	}
}
