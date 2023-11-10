package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	sRune := []rune(s)
	for i := 0; i < len(sRune); i++ {
		z01.PrintRune(sRune[i])
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			content, err := ioutil.ReadFile(args[i])
			if err != nil {
				errFile := "ERROR: open " + string(args[i]) + ": no such file or directory"
				printStr(errFile)
				z01.PrintRune('\n')
				os.Exit(1)

			}
			printStr(string(content))
		}
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	}
}
