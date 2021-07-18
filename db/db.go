package db

import "fmt"

func SaveData(input string) (int, error) {
	if input == "error" {
		return -1, fmt.Errorf("Error with %s", input)
	}
	return 100, nil
}