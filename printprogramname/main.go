package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	sTable := []rune(s[0])
	for _, name := range sTable {
		if name != '.' && name != '/' {
			z01.PrintRune(name)
		}
	}
	z01.PrintRune('\n')
}
