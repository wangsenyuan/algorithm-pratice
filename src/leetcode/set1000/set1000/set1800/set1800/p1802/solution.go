package p1802

func maxValue(n int, index int, maxSum int) int {
	// sum(A) <= maxSum
	// A[index] 最大
	// A[0] = x
	// A[index] = x
	index++
	check := func(x int) bool {
		// if A[index] == x sum(A) <= maxSum
		// x - i + 1 == 1, i == x
		// A[1...i] = 1, and A[index] = x
		// x - 1 = index - i
		i := index + 1 - x
		if i > index {
			return false
		}
		var sum int
		if i >= 1 {
			sum = i - 1
			sum += (1 + x) * (index - i + 1) / 2
		} else {
			// A[1] = x - index + 1
			sum = (x - index + 1 + x) * index / 2
		}
		// A[j] = 1 and j > index
		// j - index = x - 1
		j := x - 1 + index
		if j <= n {
			sum += n - j
			sum += (1 + x) * (j - index + 1) / 2
		} else {
			// A[n] = x - index + 1
			sum += (x - (n - index) + x) * (n - index + 1) / 2
		}
		return sum-x <= maxSum
	}

	left, right := 1, maxSum+1

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}
