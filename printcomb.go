package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for n1 := 0; n1 != 8; n1++ {
		for n2 := 1; n2 != 9; n2++ {
			for n3 := 2; n3 != 10; n3++ {
				if n1 < n2 && n2 < n3 {
					if n1 != 7 && n2 != 8 && n3 != 9{
						z01.PrintRune(rune(n1 + 48))
						z01.PrintRune(rune(n2 + 48))
						z01.PrintRune(rune(n3 + 48))
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune(rune(7 + 48))
	z01.PrintRune(rune(8 + 48))
	z01.PrintRune(rune(9 + 48))
	z01.PrintRune('\n')
}

// package piscine

// import "github.com/01-edu/z01"

// func PrintComb() {
// 	const c0, c1, c2, c7, c8, c9 = 0x30, 0x31, 0x32, 0x37,0x38, 0x39
// 	end := 0
// 	var toPrint1, toPrint2, toPrint3 rune = 0x30, 0x30, 0x30
// 	debug := 0
// 	toPrint1 = 0x30
// 	toPrint2 = 0x31
// 	toPrint3 = 0x32
// 	for end != 789 {
// 		end = int((toPrint1-48)*100) + int((toPrint2-48)*10) + int((toPrint3-48)*1)
// 		if end != 789 {
// 			z01.PrintRune(toPrint1)
// 			z01.PrintRune(toPrint2)
// 			z01.PrintRune(toPrint3)
// 			z01.PrintRune(',')
// 			z01.PrintRune(' ')
// 			if int(toPrint3) == c9 {
// 				if int(toPrint2) == c8 {
// 					toPrint1 = toPrint1+1
// 					toPrint2 = toPrint1+1
// 				}
// 				toPrint2 = toPrint2+1
// 				toPrint3 = toPrint2+1
// 			} else {
// 				toPrint3++
// 			}
// 		}
// 		if debug == 120 {
// 			break
// 		}
// 		debug++
// 	}
// 	z01.PrintRune(toPrint1)
// 	z01.PrintRune(toPrint2)
// 	z01.PrintRune(toPrint3)
// 	z01.PrintRune('\n')
// }
