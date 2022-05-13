package main

import (
	"fmt"
	"sync"
	"time"
)

func printTable(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(11)
	for number := 2; number <= 12; number++ {
		go printTable(number, &wg)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Wait()
}
