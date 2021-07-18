package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := doSth()
	fmt.Println(r, err)
}
// error = code, detail

var myError = errors.New("My new error")

type MyError struct {
	code int
	detail string
}

func (me MyError) Error() string {
	return fmt.Sprintf("My error with code %v and detail %v", me.code, me.detail)
}

func doSth() (result string, err error) {
	result = "result"
	err = MyError{code: 404, detail: "XYZ"}
	return 
}