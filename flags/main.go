package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func PrintInfo() {
	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func Sort(s string) {
	var array [1000]int
	for _, char := range s {
		array[int(char)]++
	}
	for index, char := range array {
		for char > 0 {
			z01.PrintRune(rune(index))
			char--
		}
	}
	z01.PrintRune('\n')
}

func main() {
	arguments := os.Args[1:]
	str := ""
	check := false
	SortIt := false
	for _, r := range arguments {
		check = true
		if r == "-h" || r == "--help" {
			PrintInfo()
			break
		}
		len := 0
		for i := range r {
			len = i + 1
		}
		if len > 0 {
			if r[0] == '-' {
				if len > 2 && r[2] == 'i' {
					if len > 8 {
						str += r[9:]
					}
				}
				if r[1] == 'i' {
					if len > 3 {
						str += r[3:]
					}
				}
			} else {
				str = r + str
			}
		}
		if r == "-o" || r == "--order" {
			SortIt = true
		}
	}

	if !check {
		PrintInfo()
	}
	if SortIt {
		Sort(str)
	} else {
		fmt.Println(str)
	}
}
