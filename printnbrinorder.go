package piscine

import (
	"github.com/01-edu/z01"
)

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	c := []int{}

	for n != 0 {
		c = append(c, n%10)
		n = n / 10
	}
	BubbleSort(c)

	for _, r := range c {
		z01.PrintRune(rune(r + 48))
	}
}
