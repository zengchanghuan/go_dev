package main

import (
	"fmt"
)

//函数也可以作为返回值
func add2(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func do(s string) func(int, int) int {
	switch s {
	case "+":
		return add2
	case "-":
		return sub
	default:
		return nil
	}
}
func main() {
	var a = do("+")
	fmt.Println(a(10, 20))
}
