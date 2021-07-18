package main

import "fmt"

type department struct {
	depName string
}
type person struct {
	id int
	name string
	department // Embedded struct
}

func (p person) doSth() {
	p.name = "Update name"
	fmt.Println("Called doSth() with name=", p.name)
}

func main() {
    d := department{"demo"}
	p := person{id: 1, name: "somkiat", department: d}
	p.doSth()
	fmt.Println(p.depName)
	fmt.Println(p.department.depName)
}