package main

import (
	"os"

	"github.com/01-edu/z01"
)

func TrimAtoi(s string) int {
	Result := 0
	a := 1
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b := int(r - 48)
			Result = Result*10 + b
		} else if r == '-' && Result == 0 {
			a = -1
		}
	}
	return Result * a
}

func main() {
	arguments := os.Args
	if len(arguments) > 1 {
		count := 96
		if os.Args[1] == "--upper" {
			count = 64
			for i, r := range arguments {
				if i > 0 {
					c := TrimAtoi(r) + count
					if (c >= 97 && c <= 122) || (c >= 65 && c <= 90) {
						z01.PrintRune((rune(c)))
					}
				}
			}
			z01.PrintRune('\n')
		} else {
			for i, r := range arguments {
				if i > 0 {
					c := TrimAtoi(r) + count
					if (c >= 97 && c <= 122) || (c >= 65 && c <= 90) {
						z01.PrintRune((rune(c)))
					} else if ((c < 97 || c > 122) && (c < 65 || c > 90)) && c != '-' {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
