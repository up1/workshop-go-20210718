package day01

import "fmt"

func doSth(input int) func(x int) int {
	return func(x int) int {
		return input * x
	}
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
