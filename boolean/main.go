package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args
	EvenMsg, OddMsg := "I have an even number of arguments", "I have an odd number of arguments"
	if isEven(len(args) - 1) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
