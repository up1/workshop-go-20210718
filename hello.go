package main

import "fmt"

func main() {
	var result string
	result = hello("somkiat")
	fmt.Println(result)

	row, err := saveData("error")
	if err != nil {
		fmt.Println(row, err)
	}
}

func doSth(input int) func(x int) int {
	return func(x int) int {
		return input * x
	}
}

func saveData(input string) (int, error) {
	if input == "error" {
		return -1, fmt.Errorf("Error with %s", input)
	}
	return 100, nil
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
