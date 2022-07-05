package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
	// tabbleA := []byte(a)
	// tabbleB := []byte(b)
	// iMaxA := len(tabbleA) - 1
	// iMaxB := len(tabbleB) - 1
	// resultA, resultB := 0, 0
	// for i := 0; i <= iMaxA; i++ {
	// 	resultA += int(tabbleA[i])
	// }
	// for i := 0; i <= iMaxB; i++ {
	// 	resultB += int(tabbleB[i])
	// }
	// if resultA == resultB {
	// 	return 0
	// } else if resultA < resultB {
	// 	return -1
	// } else if resultA > resultB {
	// 	return 1

	// }
	// return 1
}
