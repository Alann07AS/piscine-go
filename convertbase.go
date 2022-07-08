package piscine

import "fmt"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	baseFromTable := []rune(baseFrom)
	nbrTable := []rune(nbr)
	var valueTable []int
	baseToInt := 0
	baseFromNB := len(baseFromTable)
	for _, nb := range nbrTable {
		for i := 0; i <= baseFromNB-1; i++ {
			if nb == baseFromTable[i] {
				valueTable = append(valueTable, i)
			}
		}
	}
	// p := 0
	for i := len(valueTable) - 1; i >= 0; i-- {
		baseToInt += pow(baseFromNB+1, i)
	}
	fmt.Println(baseToInt)
	return ""
}

func pow(nb, p int) int {
	return 0
}
