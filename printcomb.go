package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var r rune = 48
	for i := 48 ; i != 58 ; i++ {
		r++
		z01.PrintRune(r)
		r++
		z01.PrintRune(r)
		r++
		z01.PrintRune(r)
		z01.PrintRune(',')
		z01.PrintRune(' ')
		r = 48
	}
}

