package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("TODO")
	for i:=0; i<10; i++ {
		defer fmt.Println("TODO", i)
	}
	fmt.Println("Finish")
}