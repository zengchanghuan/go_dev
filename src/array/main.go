package main

import (
	"fmt"
)

func main() {

	//c1 := sha256.Sum256([]byte("x1"))
	//c2 := sha256.Sum256([]byte("X1"))
	//fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	/*
		array1 := [3]string{"1324fad", "fdsa232", "gasda3"}
		fmt.Println(array1)

		var array2 = [...]int{12, 2324, 32435, 353, 535, 3535}
		fmt.Println(len(array2))
	*/

	var a = 10
	b := a
	a = 20
	fmt.Println(a, b)

	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1
	arr1[0] = 11
	fmt.Println(arr1)
	fmt.Println(arr2)

	//值类型与引用类型
	//var arr1 = []int{1, 2, 3}
	//arr2 := arr1
	//arr1[0] = 11
	//fmt.Println(arr1)
	//fmt.Println(arr2)

	//var arr1 = []int{1, 2, 3, 4}
	//arr2 := arr1
	//arr2[0] = 111
	//fmt.Println(arr1)
	//fmt.Println(arr2)
}
