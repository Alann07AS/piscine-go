package piscine

func BasicAtoi(s string) int {
	b := []byte(s)
	lenth := len(b) - 1
	wlenth := lenth - 1
	var ans, weight int = 0, 1
	for i := 0; i <= lenth; i++ {
		for i2 := wlenth; i2 >= 0; i2-- {
			weight *= 10
		}
		ans += (int(b[i]) - 48) * weight
		wlenth--
		weight = 1
	}
	return ans
}
