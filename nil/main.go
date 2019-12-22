package main

import "fmt"

func main() {
	var i interface{}
	i = nil
	fmt.Printf("%b\n", i == nil)
	var err error = nil
	fmt.Printf("%b\n", err == nil)
	fmt.Printf("%b\n", hoge() == nil)
}

func hoge() error {
	return fuga()
}

type MyError struct{}

func (e *MyError) Error() string {
	return "error"
}

func fuga() *MyError {
	return nil
}
