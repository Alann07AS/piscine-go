package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	croisant := false
	for i := 0; i <= len(a)-2; i++ {
		if f(a[i], a[i+1]) < 0 {
			croisant = true
			break
		}
		if f(a[i], a[i+1]) > 0 {
			croisant = false
			break
		}
	}
	isModified := true
	if croisant {
		for isModified {
			isModified = false
			for i := 0; i <= len(a)-2; i++ {
				if f(a[i], a[i+1]) < 0 {
					b := a[i]
					a[i] = a[i+1]
					a[i+1] = b
					isModified = true
				}
			}
		}
	} else {
		for isModified {
			isModified = false
			for i := 0; i <= len(a)-2; i++ {
				if f(a[i], a[i+1]) > 0 {
					b := a[i]
					a[i] = a[i+1]
					a[i+1] = b
					isModified = true
				}
			}
		}
	}
}
