package utils

// GCD returns the greatest common divisor of two integers.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of two integers.
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// IsPrime checks whether a number is prime.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// AbsDiff returns the absolute value difference between two integers.
func AbsDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}
