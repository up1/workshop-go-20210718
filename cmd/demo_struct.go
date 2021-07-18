package main

import "fmt"


type person struct {
	id int
	name string
}


func main() {
	p := person{id: 1, name: "somkiat"}
	fmt.Printf("%+v", p)
}