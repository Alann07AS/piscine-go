package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	if int64(piscine.Atoi(args[1])) > 9223372036854775807 || int64(piscine.Atoi(args[3])) < -9223372036854775808 {
		return
	}
	result := 0
	switch args[2] {
	case "*":
		result = Atoi(args[1]) * piscine.Atoi(args[3])
	case "+":
		result = Atoi(args[1]) + piscine.Atoi(args[3])
	case "-":
		result = Atoi(args[1]) - piscine.Atoi(args[3])
	case "/":
		if args[3] != "0" {
			result = Atoi(args[1]) / piscine.Atoi(args[3])
		} else {
			printStr("No division by 0")
			return
		}
	case "%":
		if args[3] != "0" {
			result = Atoi(args[1]) % piscine.Atoi(args[3])
		} else {
			printStr("No modulo by 0")
			return
		}
	}

	printStr(intToString(result))
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	var ans, weight int = 0, 1
	if !(s == "") {
		b := []byte(s)
		lenth := len(b) - 1
		wlenth := lenth - 1
		valid := true
		loop := 0
		for i := 0; i <= lenth; i++ {
			if !(b[i] >= '0' && b[i] <= '9') {
				valid = false
				if i == 0 && (b[0] == '+' || b[0] == '-') {
					valid = true
					loop = 1
					wlenth -= 1
				}
			}
		}
		if valid {
			for loop <= lenth {
				for i2 := wlenth; i2 >= 0; i2-- {
					weight *= 10
				}
				ans += (int(b[loop]) - 48) * weight
				wlenth--
				weight = 1
				loop++
			}
		} else {
			ans = 0
		}
		if b[0] == '-' {
			ans *= -1
		}
	} else {
		ans = 0
	}
	return ans
}

func intToString(nb int) string {
	str := ""
	l := 0
	nbTemp := nb
	weight := 1
	for nbTemp != 0 {
		l++
		nbTemp /= 10
	}
	for i := 1; i < l; i++ {
		weight *= 10
	}
	for i := 1; i <= l; i++ {
		str += string(rune(((nb / weight) % 10) + 48))
		weight /= 10
	}
	return str
}
