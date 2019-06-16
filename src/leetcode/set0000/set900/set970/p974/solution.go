package p974

func subarraysDivByK(A []int, K int) int {
	n := len(A)
	rem := make([]int, K)
	rem[0] = 1
	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
		for sum < 0 {
			sum += K
		}
		sum %= K
		rem[sum]++
	}

	var res int
	for i := 0; i < K; i++ {
		res += rem[i] * (rem[i] - 1) / 2
	}
	return res
}
