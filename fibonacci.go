package piscine

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n > 0 {
		return Fibonacci(n-1) + Fibonacci(n-2)
	} else {
		return -1
	}
}
