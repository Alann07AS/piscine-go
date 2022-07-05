package piscine

func TrimAtoi(s string) int {
	if OneNumMin(s) {
		if isNeg(s) {
			return onlyNumeric(s) * -1
		} else {
			return onlyNumeric(s)
		}
	} else {
		return 0
	}
}

func isNeg(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if OneNumMin(string(tabble[i])) {
			return false
		}
		if tabble[i] == '-' {
			return true
		}
	}
	return true
}

func onlyNumeric(s string) int {
	newTable := []rune(s)
	tabble := []rune(s)
	l := len(tabble) - 1
	i2 := 0
	for i := 0; i <= l; i++ {
		if IsNumeric(string(tabble[i])) {
			newTable[i2] = rune(tabble[i])
			i2++
		}
	}
	return AtoiForTrimatoi(string(newTable), i2-1)
}

func AtoiForTrimatoi(s string, l int) int {
	var ans, weight int = 0, 1
	if !(s == "") {
		b := []byte(s)
		lenth := l
		wlenth := lenth - 1
		valid := true
		loop := 0
		for i := 0; i <= lenth; i++ {
			if !(b[i] >= '0' && b[i] <= '9') {
				valid = false
				if i == 0 && (b[0] == '+' || b[0] == '-') {
					valid = true
					loop = 1
					wlenth -= 1
				}
			}
		}
		if valid {
			for loop <= lenth {
				for i2 := wlenth; i2 >= 0; i2-- {
					weight *= 10
				}
				ans += (int(b[loop]) - 48) * weight
				wlenth--
				weight = 1
				loop++
			}
		} else {
			ans = 0
		}
		if b[0] == '-' {
			ans *= -1
		}
	} else {
		ans = 0
	}
	return ans
}

func OneNumMin(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if tabble[i] >= '0' && tabble[i] <= '9' {
			return true
		}
	}
	return false
}
