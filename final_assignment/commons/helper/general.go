package helper

import (
	"fmt"
)

func RecoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
