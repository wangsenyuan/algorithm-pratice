package p927

const MOD = 1000000007

func threeEqualParts(A []int) []int {
	f := make(map[int64]int)
	n := len(A)
	var prod int64
	for i := 0; i < n; i++ {
		prod = prod*2 + int64(A[i])
		prod %= MOD
		if _, found := f[prod]; !found {
			f[prod] = i
		}
	}
	prod = 0
	factor := int64(1)
	for i := n - 1; i >= 0; i-- {
		prod = factor*int64(A[i]) + prod
		prod %= MOD
		factor = (factor * 2) % MOD
		if j, found := f[prod]; found {
			// end at j
			//the need to check elements [j+1, i)
			k := j + 1
			var tmp int64
			for k < i {
				tmp = (tmp*2 + int64(A[k])) % MOD
				k++
			}
			if tmp == prod {
				return []int{j, i}
			}
		}
	}
	return []int{-1, -1}
}
