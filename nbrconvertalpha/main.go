package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	for i := 1; i <= len(s)-1; i++ {
		ascii := string(rune(piscine.Atoi(s[i]) + 96))
		leter := ascii[0]
		if piscine.IsAlpha(ascii) {
			z01.PrintRune(rune(leter))
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
