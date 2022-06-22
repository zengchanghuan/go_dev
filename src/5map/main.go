package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}

	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}
