package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	for y:= 1; y <=8; y++ {
		z01.PrintRune(rune('a'+y-1))

	}
}

func recursiveIsCheck(y, nPion int) int, bool{

}