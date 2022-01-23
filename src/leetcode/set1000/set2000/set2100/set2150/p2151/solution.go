package p2151

func maximumGood(statements [][]int) int {
	n := len(statements)

	valid := func(mask int) bool {
		// 0 for bad, 1 for good
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 0 {
				continue
			}
			// i is good person
			for j := 0; j < n; j++ {
				if i == j || statements[i][j] == 2 {
					continue
				}
				tmp := statements[i][j]
				if tmp != (mask>>j)&1 {
					return false
				}
			}
		}
		return true
	}

	var ans int

	N := 1 << n

	for mask := 1; mask < N; mask++ {

		if valid(mask) {
			ans = max(ans, digit(mask))
		}
	}

	return ans
}

func digit(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
