package main

import "fmt"

func main() {
    var result string
	result = hello("somkiat")
	fmt.Println(result)

	row, err := saveData("error");
	fmt.Println(row, err)
	if err != nil {
	}
}

func saveData(input string) (int, error){
	if input == "error" {
		return -1, fmt.Errorf("Error with %s", input)
	}
	return 100, nil
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

