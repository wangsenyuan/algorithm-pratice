package productsum

const MOD = 1000000007

func ProductSum(a []int, m int) int {
	// your code here
	n := len(a)

	F := make([]int64, m+1)
	F[0] = 1

	for i := 1; i <= n; i++ {
		cur := int64(a[i-1])
		for j := min(i, m); j > 0; j-- {
			tmp := (cur * F[j-1]) % MOD
			F[j] += tmp
			F[j] %= MOD
		}
	}

	return int(F[m])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
