package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5} // Array
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
	n = append(n, 6)
	for i, v := range n {
		fmt.Println(i, v)
	}

	i := 0
	for i < len(n) {
		fmt.Println(n[i])
		i++
	}

	// for {
	// }
}
