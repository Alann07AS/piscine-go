package piscine

import (
	"fmt"
)

func PrintNbrBase(nbr int, base string) {
	result := 0
	weight := 0
	tabble := [40]rune{}
	var finaltable, lastTabble = []rune{}, []rune{}
	baseInt := len(base)
	for i := 0; i <= baseInt-1 && result != nbr; i++ {
		tabble[0]++
		result++
		lastTabble = recursiveCalc(baseInt, weight, tabble[:])
		if result == nbr {
			finaltable = lastTabble
		}
		fmt.Println(result) // nombre de tour dans ma boucle
	}
	fmt.Println((finaltable))
}

func recursiveCalc(base, weight int, tabble []rune) []rune {
	if tabble[weight] == rune(base) {
		recursiveCalc(base, weight+1, tabble)
	} else {
		tabble[weight]++
		return tabble
	}
	return tabble
}
