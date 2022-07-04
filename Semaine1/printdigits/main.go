package main

import "github.com/01-edu/z01"

func main() {
	var r rune = 47
	for i := 0; i != 10; i++ {
		r++
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
