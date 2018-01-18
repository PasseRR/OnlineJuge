package page1

func myPow(x float64, n int) float64 {
	result, isNegative := float64(1), false

	if n < 0 {
		isNegative, n = true, -n
	}

	for i := n; i != 0; i >>= 1 {
		if i%2 != 0 {
			result *= x
		}
		x *= x
	}

	if isNegative {
		result = 1 / result
	}

	return result
}
