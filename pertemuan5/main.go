package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go func() {
		fmt.Println("Hello1")
		data := "Hello"

		c <- data
	}()

	go func() {
		fmt.Println("Hai1")
		diti := "hai"

		c <- diti

	}()

	time.Sleep(time.Second * 2)

	result := <-c //this is blocker

	fmt.Println(result)

	close(c)

}
