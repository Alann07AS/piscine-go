package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	bytes, err1 := ioutil.ReadAll(os.Stdin)
	if err1 == nil {
		printStr(string(bytes))
		return
	}
	for i := 1; i <= len(args)-1; i++ {
		file, err := os.Open(args[i])
		if err != nil {
			printStr("ERROR: ")
			printStr(err.Error())
			os.Exit(1)
			z01.PrintRune('\n')
			return
		}
		stat, _ := file.Stat()
		contennet := make([]byte, stat.Size())
		file.Read(contennet)
		printStr(string(contennet))
		file.Close()
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
