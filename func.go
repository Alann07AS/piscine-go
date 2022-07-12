package piscine

import "fmt"

func ForEach(f func(int), a []int) {
	for _, each := range a {
		fmt.Print(each)
	}
}
