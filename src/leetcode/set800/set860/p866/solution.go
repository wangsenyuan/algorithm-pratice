package p866

import "math"

func primePalindrome(N int) int {
	for {
		if N > 1 && N == reverse(N) && isPrime(N) {
			return N
		}
		N++
		if 10000000 < N && N < 100000000 {
			N = 100000000
		}
	}
}

func reverse(N int) int {
	var res int

	for N > 0 {
		res = res*10 + N%10
		N /= 10
	}
	return res
}

func isPrime(N int) bool {
	sq := int(math.Sqrt(float64(N)))
	for x := 2; x <= sq; x++ {
		if N%x == 0 {
			return false
		}
	}
	return true
}
