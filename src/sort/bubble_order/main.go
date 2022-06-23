package main

import "fmt"

func main() {
	//冒泡排序从小到大
	/*
		var numSlice = []int{9, 6, 5, 4, 8}
		for i := 0; i < len(numSlice); i++ {
			for j := 0; j < len(numSlice)-1-i; j++ {
				if numSlice[j] > numSlice[j+1] {
					temp := numSlice[j]
					numSlice[j] = numSlice[j+1]
					numSlice[j+1] = temp
				}

			}
		}

		fmt.Println(numSlice)
	*/

	//冒泡排序从大到下
	var numSlice = []int{9, 6, 5, 4, 8}
	for i := 0; i < len(numSlice); i++ {
		for j := 0; j < len(numSlice)-1-i; j++ {
			if numSlice[j] < numSlice[j+1] {
				temp := numSlice[j]
				numSlice[j] = numSlice[j+1]
				numSlice[j+1] = temp
			}

		}
	}
	fmt.Println(numSlice)
}
