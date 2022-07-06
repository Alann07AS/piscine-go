package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	isModifie := true
	for isModifie {
		isModifie = false
		for i := len(s) - 1; i >= 1; i-- {
			a := ""
			if i >= 2 && s[i-1] > s[i] {
				a = s[i]
				s[i] = s[i-1]
				s[i-1] = a
				isModifie = true
			}
		}
	}

	for i := 1; i <= len(s)-1; i++ {
		sTable := []rune(s[i])
		for _, name := range sTable {
			z01.PrintRune(name)
		}
		z01.PrintRune('\n')
	}
}
