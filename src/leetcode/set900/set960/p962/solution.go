package p962

func maxWidthRamp(A []int) int {
	n := len(A)

	if n <= 1 {
		return 0
	}

	stack := make([]int, n)
	var p int
	var ans int
	for i := 0; i < n; i++ {
		if p > 0 {
			left, right := 0, p
			for left < right {
				mid := (left + right) / 2
				if A[stack[mid]] <= A[i] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			//A[left] <= A[i]
			if left < p {
				ans = max(ans, i-stack[left])
			}
		}

		if p == 0 || A[stack[p-1]] > A[i] {
			stack[p] = i
			p++
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
