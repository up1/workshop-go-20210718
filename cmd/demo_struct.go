package main

import (
	"day01"
	"fmt"
)

func main() {
	p := day01.NewPerson(1, "somkiat", "demo")
	p.DoSth()
	fmt.Printf("%+v", p)
	fmt.Println(p.DepName)
	fmt.Println(p.Department.DepName)
}
