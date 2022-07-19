package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Fprime(value int) {
	if value != 1 {
		divisionIterator := 2
		for value > 1 {
			if value%divisionIterator == 0 {
				fmt.Print(divisionIterator)
				value = value / divisionIterator
				if value > 1 {
					z01.PrintRune('*')
				}
				divisionIterator--
			}
			divisionIterator++
		}
	}
	z01.PrintRune('\n')
}
