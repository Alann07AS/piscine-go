package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}
	stat, _ := file.Stat()
	contennet := make([]byte, stat.Size())
	file.Read(contennet)
	fmt.Print(string(contennet))
}
