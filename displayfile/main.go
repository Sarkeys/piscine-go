package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "quest8.txt"
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf(string(file))
}
