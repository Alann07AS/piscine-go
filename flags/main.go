package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	insertFilter := []string{"--insert=", "-i="}
	orderFilter := []string{"--order", "-o"}
	isInsert := false
	isOrder := false
	insertPos := 0
	argPos := 1
	argumentTable := os.Args
	if len(argumentTable) == 1 || argumentTable[1] == "-h" || argumentTable[1] == "--help" {
		printHelp()
		return
	}
	for i, flags := range argumentTable {
		table := []rune(flags)
		l := len(table) - 1
		equalPos := 0
		for i := 0; i <= l; i++ {
			if table[i] == '=' {
				equalPos = i + 1
				i = l
			}
		}
		sToCheckInsert := string(table[:equalPos])
		if sToCheckInsert == insertFilter[0] || sToCheckInsert == insertFilter[1] {
			isInsert = true
			insertPos = i
		}
		if flags == orderFilter[0] || flags == orderFilter[1] {
			isOrder = true
		}
		if !(flags == insertFilter[0] || flags == insertFilter[1]) && !(flags == orderFilter[0] || flags == orderFilter[1]) && i != 0 {
			argPos = i
		}
	}
	if isInsert {
		table := []rune(argumentTable[insertPos])
		l := len(table) - 1
		sToConcat := ""
		firstCharPos := 0
		for i := 0; i <= l; i++ {
			if table[i] == '=' {
				firstCharPos = i + 1
				i = l
			}
		}
		sToConcat = string(table[firstCharPos : l+1])
		argumentTable[argPos] = concatArgs(argumentTable[argPos], sToConcat)
		if !isOrder {
			printTxt(argumentTable[argPos])
		}
	}
	if isOrder {
		isModifie := true
		s := []rune(argumentTable[argPos])
		for isModifie {
			isModifie = false
			for i := len(s) - 1; i >= 1; i-- {
				var a rune
				if i >= 1 && s[i-1] > s[i] {
					a = s[i]
					s[i] = s[i-1]
					s[i-1] = a
					isModifie = true
				}
			}
		}
		printTxt(string(s))
	}
	if !isInsert && !isOrder {
		printTxt(argumentTable[argPos])
	}
	z01.PrintRune('\n')
}

func concatArgs(str1 string, str2 string) string {
	return str1 + str2
}

func printHelp() {
	help := []string{
		"--insert\n",
		"  -i\n",
		"         This flag inserts the string into the string passed as argument.\n",
		"--order\n",
		"  -o\n",
		"         This flag will behave like a boolean, if it is called it will order the argument.\n",
	}
	for _, txt := range help {
		printTxt(txt)
	}
}

func printTxt(s string) {
	table := []rune(s)
	for _, leter := range table {
		z01.PrintRune(leter)
	}
}
