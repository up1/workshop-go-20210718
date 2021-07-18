package main

import (
	"day01"
	"fmt"
)


type parent struct{
}

type child struct {
	parent
}

func (p parent) callParent() {
	fmt.Println("callParent..")
}

// func (c child) callParent() {
// 	c.parent.callParent() // Call from parent
// 	fmt.Println("callParent from child..")
// }

func main() {
	c := child{}
	c.callParent()

	p := day01.NewPerson(1, "somkiat", "demo")
	p.DoSth()
	fmt.Printf("%+v", p)
	fmt.Println(p.DepName)
	fmt.Println(p.Department.DepName)
}
