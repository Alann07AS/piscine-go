package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	const z = 48
	var d1, d2, d3, d4 rune
	for c := 1; c != 9900; c++ {
		d1 = rune(c/1000 + z)
		d2 = rune((c/100)%10 + z)
		d3 = rune((c/10)%10 + z)
		d4 = rune(c%10 + z)
		if int(d1)-z != int(d3)-z && int(d2)-z != int(d4)-z {
			z01.PrintRune(d1)
			z01.PrintRune(d2)
			z01.PrintRune(' ')
			z01.PrintRune(d3)
			z01.PrintRune(d4)
			if !(int(d1)-z == 9 && int(d2)-z == 8 && int(d3)-z == 9 && int(d4)-z == 9) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
