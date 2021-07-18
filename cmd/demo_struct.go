package main

import (
	"day01"
	"fmt"
)

func main() {
	d := day01.Department{"demo"}
	p := day01.Person{
		Id:         1,
		Name:       "somkiat",
		Department: d}
	p.DoSth()
	fmt.Println(p.DepName)
	fmt.Println(p.Department.DepName)
}
