package piscine

func SplitWhiteSpaces(s string) []string {
	table := []rune(s)
	var sTable []string
	indexD := 0
	l := 17
	for i := 0; i <= l; i++ {
		if table[i] == ' ' || table[i] == '	' || table[i] == '\n' || i == l {
			if i == l {
				sTable = append(sTable, string(table[indexD:l+1]))
			} else {
				sTable = append(sTable, string(table[indexD:i]))
				indexD = i + 1
			}
		}
	}
	return sTable
}
