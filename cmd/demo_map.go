package main

import "fmt"

func main() {

	n := 7
	switch n {
		case 1,5,6,7 : fmt.Println("OK 1")
		case 2 : fmt.Println("OK 2")
		case 3 : fmt.Println("OK 3")
		default : fmt.Println("default")
	}

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