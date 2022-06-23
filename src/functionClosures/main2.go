package main

import "fmt"

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	var f = adder2(10)
	fmt.Println(f(10))
	fmt.Println(f(20)) //40
}
