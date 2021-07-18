package main

import "fmt"

func main() {
	grades := make(map[int]string)
	grades[100] = "A"
	grades[101] = "B"
	grades[102] = "C+"

	for k, v := range grades {
		fmt.Println(k, v)
	}

	fmt.Printf("Data of 100=%v\n", grades[100])
	
	if v, found := grades[10111]; found {
		fmt.Println("Found=", v)
	} else {
		fmt.Println("Not found", v)
	}

}