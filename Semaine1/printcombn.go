package piscine

import "github.com/01-edu/z01"

// import "fmt"

func PrintCombN(n int) {
	// fmt.Println("INIT")
	comb := 0
	if n == 9 {
		comb = 12345678
	}
	lastComb := LastComb(n)
	for comb <= lastComb {
		if combIs(comb, n) && comb != lastComb {
			Pcomb(comb, n)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		if comb == lastComb {
			Pcomb(comb, n)
			z01.PrintRune('\n')
			// fmt.Println("Is Last Comb")
			break
		}
		if comb%10000000 == 3456789 {
			comb += 10000000
		} else {
			comb++
		}
	}
}

func LastComb(n int) int {
	lastComb := 9
	weight := 10
	for i := 1; i < n; i++ {
		lastComb = lastComb + ((9 - i) * weight)
		weight *= 10
	}
	// fmt.Println("Lastcomb is ", lastComb)
	return lastComb
}

func combIs(comb, n int) bool {
	stringComb := IntToString(comb, n)
	tableComb := []byte(stringComb)
	state := true
	for i := 0; i < n-1; i++ {
		if tableComb[i] >= tableComb[i+1] {
			state = false
		}
	}
	// fmt.Println(tableComb)
	return state
}

func IntToString(nb int, n int) string {
	weight := 1
	table := []byte("::::::::::")
	for i := 1; i < n; i++ {
		weight *= 10
	}
	for i2 := 0; i2 <= n; i2++ {
		table[i2] = byte((nb/weight)%10 + '0')
		if weight != 1 {
			weight /= 10
		}
	}
	return string(table)
}

func Pcomb(comb, n int) {
	stringComb := IntToString(comb, n)
	tableComb := []byte(stringComb)
	for i2 := 0; i2 <= n-1; i2++ {
		z01.PrintRune(rune(tableComb[i2]))
	}
}
