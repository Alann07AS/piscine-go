package piscine

func ConcatParams(args []string) string {
	concatS := ""
	for i := 0; i < len(args); i++ {
		concatS += args[i]
		if i != len(args)-1 {
			concatS += "\n"
		}
	}
	return concatS
}
