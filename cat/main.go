package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := 1; i <= len(args)-1; i++ {
		file, err := os.Open(args[i])
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
			return
		}
		stat, _ := file.Stat()
		contennet := make([]byte, stat.Size())
		file.Read(contennet)
		fmt.Print(string(contennet))
		file.Close()
	}
}
