package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	result := 0
	weight := calcWeight(n)
	var table [40]int
	l := calcL(weight)
	for i := 0; i < l; i++ {
		result = (n / weight) % 10
		table[i] = result
		weight /= 10
	}
	isModifie := true
	for isModifie {
		isModifie = false
		for i := l - 1; i >= 0; i-- {
			if i >= 1 && table[i-1] > table[i] {
				table[39] = table[i]
				table[i] = table[i-1]
				table[i-1] = table[39]
				isModifie = true
			}
		}
	}
	for i := 0; i <= l-1; i++ {
		z01.PrintRune(rune(table[i] + 0x30))
	}
}

func calcWeight(n int) int {
	weight := 1
	for i := n; i/10 > 0; i /= 10 {
		weight *= 10
	}
	return weight
}

func calcL(weight int) int {
	l := 0
	for weight != 0 {
		l++
		weight /= 10
	}
	return l
}
