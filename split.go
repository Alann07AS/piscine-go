package piscine

func Split(s, sep string) []string {
	table := []rune(s)
	tableSep := []rune(sep)
	var sTable []string
	indexD := 0
	l := len(table) - 1
	isSeparate := false
	for i := 0; i <= l; i++ {
		if table[i] == tableSep[0] || i == l {
			if i == l {
				if i == l {
					sTable = append(sTable, string(table[indexD:l+1]))
				}
			} else {
				for i2 := 1; i2 <= len(tableSep)-1; i2++ {
					i++
					if table[i] == tableSep[i2] {
						isSeparate = true
					}
				}
				if isSeparate {
					sTable = append(sTable, string(table[indexD:i-len(tableSep)+1]))
					indexD = i + 1
					isSeparate = false
				}
			}
		}
	}
	return sTable
}
