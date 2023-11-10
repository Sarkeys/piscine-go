package main

import (
	"fmt"
	"os"
)

func Integer(s string) int {
	result := 0
	num := 1
	for i := len(s) - 1; i >= 0; i-- {
		n := int(s[i]) - '0'
		result += n * num
		num *= 10
	}
	return result
}

func main() {
	if len(os.Args) > 3 {
		num := Integer(os.Args[2])
		args := os.Args[3:]
		exit := 0
		count := 0
		for _, char := range args {
			file, err := os.Open(char)
			if err != nil {
				exit = 1
				fmt.Printf(err.Error())
				fmt.Printf("\n")
				count++
			} else {
				ans, _ := os.ReadFile(char)
				len := len(ans) - num
				if len < 0 {
					len = 0
				}
				if count == 0 {
					fmt.Printf("==> %v <==\n%v", char, string(ans[len:]))
					count++
				} else if count >= 1 {
					fmt.Printf("\n==> %v <==\n%v", char, string(ans[len:]))
				}
			}
			file.Close()
		}
		os.Exit(exit)
	}
}
