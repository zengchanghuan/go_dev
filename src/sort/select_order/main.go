package main

import "fmt"

func main() {
	/*
		var arr = [...]int{1, 3, 5, 7, 8}
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i]+arr[j] == 8 {
					fmt.Printf("(%v,%v)", i, j)
				}
			}
		}

	*/

	//从小到大的排序
	/*
		var numSlice = []int{9, 8, 7, 6, 5, 4}
		for i := 0; i < len(numSlice); i++ {
			for j := i + 1; j < len(numSlice); j++ {
				if numSlice[i] > numSlice[j] {
					temp := numSlice[i]
					numSlice[i] = numSlice[j]
					numSlice[j] = temp
				}

			}
		}
		fmt.Println(numSlice)
	*/

	//从大到小的排序
	var numSlice = []int{6, 5, 4, 9, 8, 7}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] < numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}

		}
	}
	fmt.Println(numSlice)
}
