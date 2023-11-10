package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Prints(x, y int) {
	var data1 []rune
	var data2 []rune
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	dec := "0123456789"
	for i := 0; i < x; i++ {
		data1 = append(data1, (rune(dec[x%len(dec)])))
		x /= len(dec)
	}
	for i := len(data1) - 1; i >= 0; i-- {
		z01.PrintRune(data1[i])
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := 0; i < y; i++ {
		data2 = append(data2, rune(dec[y%len(dec)]))
		y /= len(dec)
	}
	for i := len(data1) - 1; i >= 0; i-- {
		z01.PrintRune(data2[i])
	}
	z01.PrintRune('\n')
}

func main() {
	points := new(point)

	setPoint(points)
	Prints(points.x, points.y)
}
