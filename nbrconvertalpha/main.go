package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	for i := 1; i <= len(s)-1; i++ {
		ascii := string(rune(Atoi(s[i]) + 96))
		leter := ascii[0]
		if IsAlpha(ascii) {
			z01.PrintRune(rune(leter))
		} else {
			z01.PrintRune(' ')
		}
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

func IsAlpha(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] >= '0' && tabble[i] <= '9') && !(tabble[i] >= 'a' && tabble[i] <= 'z') && !(tabble[i] >= 'A' && tabble[i] <= 'Z') {
			return false
		}
	}
	return true
}
