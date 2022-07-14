package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := 1; i <= len(args)-1; i++ {
		switch args[i] {
		case "01", "galaxy", "galaxy 01":
			fmt.Println("Alert!!!")
		}
	}
}
