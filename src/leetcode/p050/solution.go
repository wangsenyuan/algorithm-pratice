package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	if n == 0 {
		return 1.0
	}
	y := myPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}

	return y * y * x
}
