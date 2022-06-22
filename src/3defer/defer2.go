package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("error count not create file.")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "learning Go!"); error != nil {
		fmt.Println("error cloud not write to file.")
		return
	}
	newfile.Sync()
}
