package p828

const MOD = 1000000007

func uniqueLetterString(S string) int {
	n := len(S)
	seen := make(map[byte]int)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < len(S); i++ {
		if j, found := seen[S[i]]; found {
			left[i] = j
		} else {
			left[i] = -1
		}
		seen[S[i]] = i
	}
	seen = make(map[byte]int)
	for i := n - 1; i >= 0; i-- {
		if j, found := seen[S[i]]; found {
			right[i] = j
		} else {
			right[i] = n
		}
		seen[S[i]] = i
	}

	var ans int64
	for i := 0; i < n; i++ {
		l, r := left[i], right[i]
		a, b := i-l, r-i
		ans = (ans + int64(a)*int64(b)) % MOD
	}
	return int(ans)
}
