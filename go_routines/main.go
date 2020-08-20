package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go mozart(&wg)
	go tempestade(&wg)

	wg.Wait()
}

func mozart(wg *sync.WaitGroup) {
	defer wg.Done()

	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("Mozart...")
		time.Sleep(time.Second)
	}
}

func tempestade(wg *sync.WaitGroup) {
	defer wg.Done()

	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("Tempestade...")
		time.Sleep(time.Second)
	}
}
