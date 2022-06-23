package main

import "fmt"

func main() {
	//var slice []int
	//fmt.Printf("%T", slice)

	//months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	//quarter1 := months[0:3]
	//quarter2 := months[3:6]
	//quarter3 := months[6:9]
	//quarter4 := months[9:12]
	//fmt.Println(quarter1, len(quarter1), cap(quarter1))
	//fmt.Println(quarter2, len(quarter2), cap(quarter2))
	//fmt.Println(quarter3, len(quarter3), cap(quarter3))
	//fmt.Println(quarter4, len(quarter4), cap(quarter4))

	//a := [5]int{44, 55, 324, 645}
	//b := a[:]
	//fmt.Println(b)
	//fmt.Printf("%v %T \n", b, b)
	//
	////左包、右不包
	//c := a[1:3]
	//fmt.Println(c)
	//
	//d := a[2:]
	//fmt.Println(d)

	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d cap:%d\n", a, a, len(a), cap(a))

	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d cap:%d\n", b, b, len(b), cap(b))
}
