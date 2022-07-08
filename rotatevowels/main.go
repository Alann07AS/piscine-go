package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	sTable := os.Args
	sTtoWrite := ""
	if len(sTable) >= 1 {
		sTtoWrite += sTable[1]
		for i := 2; i <= len(sTable)-1; i++ {
			sTtoWrite += " "
			sTtoWrite += sTable[i]
		}
	}
	runeTable := []rune(sTtoWrite)
	var fowelTable []rune
	for _, char := range runeTable {
		if isFowel(char) {
			fowelTable = append(fowelTable, char)
		}
	}
	i := len(fowelTable) - 1
	for _, char := range runeTable {
		if isFowel(char) {
			z01.PrintRune(fowelTable[i])
			i--
		} else {
			z01.PrintRune(char)
		}
	}
}

func isFowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
