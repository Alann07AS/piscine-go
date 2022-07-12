package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := 3; i <= len(args)-1; i++ {
		file, err := os.Open("file1.txt")
		if err != nil {
			fmt.Println(err.Error())
		}
		stat, _ := file.Stat()
		contennet := make([]byte, stat.Size())
		file.ReadAt(contennet, stat.Size())
		fmt.Println(string(contennet))
		file.Close()
	}
}
