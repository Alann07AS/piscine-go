package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
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

func main() {
	points := point{}

	setPoint(&points)

	printStr("x = ")
	printStr(intToString(points.x))
	printStr(", y = ")
	printStr(intToString(points.y))

	z01.PrintRune('\n')
}
