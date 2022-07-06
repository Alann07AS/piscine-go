package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	for i := 1; i <= len(s)-1; i++ {
		sTable := []rune(s[i])
		for _, name := range sTable {
			z01.PrintRune(name)
		}
		z01.PrintRune('\n')
	}
}
