package p795

func numSubarrayBoundedMax(A []int, L int, R int) int {
	n := len(A)

	left := make([]int, n)
	stack := make([]int, n)
	p := 0
	for i := 0; i < n; i++ {
		for p > 0 && A[i] > A[stack[p-1]] {
			p--
		}
		if p > 0 {
			left[i] = stack[p-1]
		} else {
			left[i] = -1
		}
		stack[p] = i
		p++
	}

	right := make([]int, n)
	p = 0
	for i := n - 1; i >= 0; i-- {
		for p > 0 && A[i] >= A[stack[p-1]] {
			p--
		}
		if p > 0 {
			right[i] = stack[p-1]
		} else {
			right[i] = n
		}
		stack[p] = i
		p++
	}

	var ans int

	for i := 0; i < n; i++ {
		if A[i] >= L && A[i] <= R {
			x, y := left[i], right[i]
			ans += (i - x) * (y - i)
		}
	}

	return ans
}
