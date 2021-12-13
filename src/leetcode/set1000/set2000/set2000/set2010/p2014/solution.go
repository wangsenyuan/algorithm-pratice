package p2014

const A = 26

func longestSubsequenceRepeatedK(S string, K int) string {
	n := len(S)
	// n < 8 * k
	// n / k < 8
	// L < (n + k - 1) / k <= 8
	cnt := make([]int, A)
	for i := 0; i < n; i++ {
		cnt[int(S[i]-'a')]++
	}
	can := make([]int, 0, 8)
	for i := A - 1; i >= 0; i-- {
		for x := cnt[i]; x >= K; x -= K {
			can = append(can, i)
		}
	}
	// pick permutations from can, and check
	if len(can) == 0 {
		return ""
	}

	check := func(arr []int) bool {
		// check whether arr * k is a sub seq of S
		var i int
		var k int
		for k < K && i < n {
			var j int
			for j < len(arr) && i < n {
				if int(S[i]-'a') == arr[j] {
					i++
					j++
				} else {
					i++
				}
			}
			if j < len(arr) {
				return false
			}
			k++
		}
		return k == K
	}

	for l := len(can); l > 0; l-- {
		ans := permutate(can, l, check)
		if len(ans) == 0 {
			continue
		}
		buf := make([]byte, len(ans))

		for i := 0; i < len(ans); i++ {
			buf[i] = byte(ans[i] + 'a')
		}
		return string(buf)
	}

	return ""
}

func permutate(can []int, l int, check func([]int) bool) []int {
	n := len(can)
	buf := make([]int, l)
	var loop func(flag int, i int) bool

	loop = func(flag int, i int) bool {
		if i == l {
			return check(buf)
		}
		for j := 0; j < n; j++ {
			if (flag>>j)&1 == 0 {
				buf[i] = can[j]
				if loop(flag|(1<<j), i+1) {
					return true
				}
			}
		}
		return false
	}
	if loop(0, 0) {
		return buf
	}
	return nil
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
