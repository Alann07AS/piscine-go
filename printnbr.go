package piscine

import "github.com/01-edu/z01"

func Decomposeur(n int) {
	lenght := 0
	dec := 1
	if n < 0 {
		n *= -1
	}
	nb := n
	for nb/10 != 0 {
		lenght++
		nb /= 10
	}
	for r :=0; r != lenght; r++ {
		dec *= 10
	}
	for i := 0; i <= lenght; i++ {
		z01.PrintRune(rune((n/dec)%10+48))
		dec /= 10
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		Decomposeur(n)
	} else {
		Decomposeur(n)
	}
}
