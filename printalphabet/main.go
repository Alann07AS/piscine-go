package main

import  "github.com/01-edu/z01"

func main() {
	var r rune = 96
	for i := 96; i != 122; i++ {
	r++
	z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
