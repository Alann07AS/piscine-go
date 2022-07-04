package piscine

func BasicAtoi2(s string) int {
	b := []byte(s)
	lenth := len(b) - 1
	wlenth := lenth - 1
	var ans, weight int = 0, 1
	valid := true
	for i := 0; i <= lenth; i++ {
		if !(b[i] >= '0' && b[i] <= '9') {
			valid = false
		}
	}
	if valid {
		for i := 0; i <= lenth; i++ {
			for i2 := wlenth; i2 >= 0; i2-- {
				weight *= 10
			}
			ans += (int(b[i]) - 48) * weight
			wlenth--
			weight = 1
		}
	} else {
		ans = 0
	}
	return ans
}
