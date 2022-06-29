package main

import "github.com/01-edu/z01"

func main() {
	var r rune = 123
	for i := 0; i != 26; i++ {
		r--
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
