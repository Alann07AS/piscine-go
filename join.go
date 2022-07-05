package piscine

func Join(strs []string, sep string) string {
	iMax := len(strs) - 1
	cacheString := strs[0]
	for i := 0; i <= iMax-1; i++ {
		cacheString += sep
		cacheString += strs[i+1]
	}
	return cacheString
}
