package main

import (
	"fmt"
	"os"
	"piscine"
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
		result = piscine.Atoi(args[1]) * piscine.Atoi(args[3])
	case "+":
		result = piscine.Atoi(args[1]) + piscine.Atoi(args[3])
	case "-":
		result = piscine.Atoi(args[1]) - piscine.Atoi(args[3])
	case "/":
		if args[3] != "0" {
			result = piscine.Atoi(args[1]) / piscine.Atoi(args[3])
		} else {
			fmt.Println("No division by 0")
			return
		}
	case "%":
		if args[3] != "0" {
			result = piscine.Atoi(args[1]) % piscine.Atoi(args[3])
		} else {
			fmt.Println("No modulo by 0")
			return
		}
	}
	fmt.Println(result)
}
