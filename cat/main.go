package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i <= len(args)-1; i++ {
		file, err := os.Open(args[i])
		if err != nil {
			printStr("ERROR: ")
			printStr(err.Error())
			z01.PrintRune('\n')
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
	stringIN, _ := ioutil.ReadAll(os.Stdin)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode()&os.ModeCharDevice) == 0 && len(args) == 1 {
		printStr(string(stringIN))
		return
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
