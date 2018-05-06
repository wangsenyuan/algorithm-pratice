package p829

import "math"

func consecutiveNumbersSum(N int) int {

	// 2 * a0  + n - 1 = 2 * N / n

	// max n = sqrt(2 * N)
	n := int(math.Sqrt(float64(2 * N)))

	ans := 1

	for i := 2; i <= n; i++ {
		if (2*N)%i == 0 {
			x := 2 * N / i
			y := x + 1 - i
			if y%2 == 0 {
				ans++
			}
		}
	}

	return ans
}
