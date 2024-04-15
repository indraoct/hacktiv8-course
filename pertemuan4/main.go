package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go printData(i)
	}

	wg.Wait()
}

func printData(data int) {
	fmt.Println("ngeprint: ", data)
	wg.Done()
}
