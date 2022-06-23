package main

import "fmt"

func main() {
	/*
		studentsAge := make(map[string]int)
		studentsAge["john"] = 32
		studentsAge["bob"] = 31
		for name, age := range studentsAge {
			fmt.Printf("%s\t%d\n", name, age)
		}

		for name := range studentsAge {
			fmt.Printf("Names %s\n", name)
		}
	*/

	/*
		var userinfo = map[string]string{
			"username": "张三",
			"age":      "20",
			"sex":      "男",
		}

		for k, v := range userinfo {
			fmt.Printf("key:%v value:%v\n", k, v)
		}
	*/

	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果 key 存在 ok 为 true,v 为对应的值;不存在 ok 为 false,v 为值类型的零值
	v, ok := scoreMap["张四"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
