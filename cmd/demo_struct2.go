package main

import "fmt"

type processor interface {
	doSth()
}

type demo struct {
	id int
	name string
}

func (d demo) doSth() {
	fmt.Println("Called doSth")
}

func (d demo) String() string {
	return fmt.Sprintf("id=%v, name=%v", d.id, d.name)
}

type A int
func (a A) doSth() {
	fmt.Println("Called doSth from A")
}

func main() {
  do(demo{})
  do(A(1))
}

func do(p processor) {
	p.doSth()
}