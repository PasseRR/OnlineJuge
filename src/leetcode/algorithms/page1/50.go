package page1

func myPow(x float64, n int) float64 {
	result := float64(1)

	// 负n次方
	if n < 0 {
		x, n = 1/x, -n
	}

	// x^2n = x^n * x^n
	// x^(2n + 1) = x^n * x^n * x
	// 2^10 = 2^4 * 2^4 * 2 * 2
	for i := n; i != 0; i >>= 1 {
		if i%2 != 0 {
			result *= x
		}
		x *= x
	}

	return result
}
