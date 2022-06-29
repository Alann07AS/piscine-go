package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	const c0, c1, c2, c8, c9 = 0x30, 0x31, 0x32, 0x38, 0x39
	end := 0
	var toPrint1, toPrint2, toPrint3 rune = 0x30, 0x30, 0x30
	toPrint1 = 0x30
	toPrint2 = 0x31
	toPrint3 = 0x32
	for end != 789 {
		end = int((toPrint1-48)*100) + int((toPrint2-48)*10) + int((toPrint3-48)*1)
		if end != 789 {
			z01.PrintRune(toPrint1)
			z01.PrintRune(toPrint2)
			z01.PrintRune(toPrint3)
			z01.PrintRune(',')
			z01.PrintRune(' ')
			if int(toPrint3) >= c9 {
				toPrint2 = (1 + toPrint2)
				toPrint3 = (toPrint2 + 1)
			} else {
				if int(toPrint1) < int(toPrint3) && int(toPrint2) < int(toPrint3) {
					toPrint3++
				}
			}
			if int(toPrint2) >= c8 {
				toPrint1 = (1 + toPrint1)
				toPrint2 = (toPrint1 + 1)
			}
			if int(toPrint1) >= c9 {
				toPrint3 = c0
			}
			if int(toPrint3) > c9 {
				toPrint2 = (1 + toPrint2)
				toPrint3 = (toPrint2 + 1)
			}
		}
	}
	z01.PrintRune(toPrint1)
	z01.PrintRune(toPrint2)
	z01.PrintRune(toPrint3)
	z01.PrintRune('\n')
}
