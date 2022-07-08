package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	baseTableBefore := []rune(base)
	var intTbale []int
	baseNb := len(baseTableBefore) - 1
	index := 0
	isNeg := false
	if !sortAndCheckBase(baseTableBefore) || baseNb <= 0 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	baseTable := []rune(base)
	if nbr < 0 {
		nbr *= -1
		isNeg = true
	}
	for calc := 0; calc < nbr; calc++ {
		intTbale = recursifBaseCalc(baseNb, index, intTbale)
	}
	if isNeg {
		z01.PrintRune('-')
	}
	for i := len(intTbale) - 1; i >= 0; i-- {
		z01.PrintRune(baseTable[intTbale[i]])
	}
}

func recursifBaseCalc(baseNb, index int, inTable []int) []int {
	l := len(inTable) - 1
	if index > l {
		inTable = append(inTable, 0)
	}
	if inTable[index] != baseNb {
		inTable[index]++
		return inTable
	} else {
		inTable[index] = 0
		return recursifBaseCalc(baseNb, index+1, inTable)
	}
}

func sortAndCheckBase(sBase []rune) bool {
	var s []rune
	s = sBase
	isModifie := true
	for isModifie {
		isModifie = false
		for i := len(s) - 1; i >= 1; i-- {
			var a rune
			if i >= 1 && s[i-1] > s[i] {
				a = s[i]
				s[i] = s[i-1]
				s[i-1] = a
				isModifie = true
			}
		}
	}
	for i := len(s) - 1; i >= 1; i-- {
		if i >= 1 && s[i-1] == s[i] {
			return false
		}
	}
	return true
}
