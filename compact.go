package piscine

func Compact(ptr *[]string) int {
	var newTable []string
	cpt := 0
	for _, each := range *ptr {
		if each != "" {
			newTable = append(newTable, each)
			cpt++
		}
	}
	*ptr = newTable
	return cpt
}
