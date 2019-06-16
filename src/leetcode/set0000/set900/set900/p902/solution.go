package p902

import (
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	S := strconv.Itoa(N)
	K := len(S)

	dp := 1
	for i := K - 1; i >= 0; i-- {
		prevDp := dp
		var cur int
		num := int(S[i] - '0')
		for _, d := range D {
			x := int(d[0] - '0')
			if x < num {
				cur += pow(len(D), K-i-1)
			} else if x == num {
				cur += prevDp
			}
		}
		dp = cur
	}
	for i := 1; i < K; i++ {
		dp += pow(len(D), i)
	}
	return dp
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res *= a
		}
		a *= a
		n >>= 1
	}
	return res
}
