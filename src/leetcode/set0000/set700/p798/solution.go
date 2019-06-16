package p798

func bestRotation(A []int) int {
	n := len(A)
	bad := make([]int, n)

	for i := 0; i < n; i++ {
		left := (i - A[i] + 1 + n) % n
		right := (i + 1) % n
		bad[left]--
		bad[right]++
		if left > right {
			bad[0]--
		}
	}

	best := -n
	var ans int
	var cur int

	for i := 0; i < n; i++ {
		cur += bad[i]
		if cur > best {
			best = cur
			ans = i
		}
	}
	return ans
}
