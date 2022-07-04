package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i < 'z'; i+=2 {
		z01.PrintRune(i)
		z01.PrintRune(i - 31)
	}
	z01.PrintRune('\n')
}
