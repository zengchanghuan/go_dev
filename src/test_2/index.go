package main

import (
	"fmt"
)

func main() {
	/*
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println("sum 1...100 is", sum)
	*/
	fmt.Println("golang is end ********")

	//只要 num 变量保存的值与 5 不同，程序就会输出一个随机数。
	/*
		var num int64
		rand.Seed(time.Now().Unix())
		for num != 5 {
			num = rand.Int63n(15)
			fmt.Println(num)
		}
	*/
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)

}
