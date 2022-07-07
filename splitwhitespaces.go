package piscine

func SplitWhiteSpaces(s string) []string {
	table := []rune(s)
	var sTable []string
	indexD := 0
	l := len(table) - 1
	for i := 0; i <= l; i++ {
		ti := table[i]
		if ti == ' ' || ti == '	' || ti == '\n' || i == l {
			if string(table[indexD:i]) != "" {
				if i == l {
					sTable = append(sTable, string(table[indexD:l+1]))
				} else {
					sTable = append(sTable, string(table[indexD:i]))
					indexD = i + 1
				}
			} else {
				indexD++
			}
		}
	}
	return sTable
}
