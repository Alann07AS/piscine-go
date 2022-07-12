package piscine

func SortWordArr(a []string) {
	isModified := true
	for isModified {
		isModified = false
		for i := 0; i <= len(a)-2; i++ {
			if a[i] > a[i+1] {
				b := a[i]
				a[i] = a[i+1]
				a[i+1] = b
				isModified = true
			}
		}
	}
}
