package main

import (
	"day01"
	xyz "day01/db"
	"fmt"
)

func main() {
	var result string
	result = day01.Hello("somkiat")
	fmt.Println(result)

	row, err := xyz.SaveData("error")
	if err != nil {
		fmt.Println(row, err)
	}
}

