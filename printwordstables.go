package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(args []string) {
	concatS := ""
	for i := 0; i < len(args); i++ {
		concatS += args[i]
		if i != len(args)-1 {
			concatS += "\n"
		}
	}
	printTxt(concatS)
	z01.PrintRune('\n')
}

func printTxt(s string) {
	table := []rune(s)
	for _, leter := range table {
		z01.PrintRune(leter)
	}
}
