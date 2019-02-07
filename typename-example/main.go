package main

import "fmt"

type GoSheffield int

//go:generate typename GoSheffield

func main() {
	var gopher GoSheffield = 42
	fmt.Printf("Today's special number is %v (%v)\n", gopher, gopher.TypeName())
}
