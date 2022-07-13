package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := 3; i <= len(args)-1; i++ {
		file, err := os.Open(args[i])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			stat, _ := os.Stat(file.Name())
			contennet := make([]byte, stat.Size())
			lookAt := Atoi(args[2])
			file.ReadAt(contennet, stat.Size()-int64(lookAt))
			fmt.Println(string(contennet))
			file.Close()
		}
	}
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
