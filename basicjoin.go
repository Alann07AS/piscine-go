package piscine

func BasicJoin(elems []string) string {
	iMax := len(elems) - 1
	cacheString := elems[0]
	for i := 0; i <= iMax-1; i++ {
		cacheString += elems[i+1]
	}
	return cacheString
}
