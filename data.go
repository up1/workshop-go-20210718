package day01

import "fmt"

type Department struct {
	DepName string
}
type Person struct {
	Id         int
	Name       string
	Department // Embedded struct
}

func (p Person) DoSth() {
	p.Name = "Update name"
	fmt.Println("Called doSth() with name=", p.Name)
}
