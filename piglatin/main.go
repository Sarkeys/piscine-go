package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	first := ""
	second := ""
	isvowel := false
	for i, char := range str {
		if !IsVowel(byte(str[0])) && IsVowel(byte(char)) {
			first = str[:i]
			second = str[i:]
			isvowel = true
		}
	}
	if isvowel {
		fmt.Print(second, first, "ay", "\n")
	} else if IsVowel(byte(str[0])) {
		fmt.Print(str, "ay", "\n")
	} else {
		fmt.Println("No vowels")
	}
}

func IsVowel(iscoman byte) bool {
	vowels := []byte{'A', 'E', 'I', 'O', 'U', 'Y', 'a', 'e', 'i', 'o', 'u', 'y'}
	for _, char := range vowels {
		if char == iscoman {
			return true
		}
	}
	return false
}
