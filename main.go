package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println("I am happy")

	for i:=0; i<10; i++{
		if !(i >= 10) {
			fmt.Println("Counter is i", i)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("You have reached i", i)

	}
}
