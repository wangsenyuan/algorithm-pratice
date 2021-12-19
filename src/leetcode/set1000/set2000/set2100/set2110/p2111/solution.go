package p2111

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	stack := make([]int, n)
	buf := make([]int, n)
	var res int
	for i := 0; i < k; i++ {
		var p int
		for j := i; j < n; j += k {
			buf[p] = arr[j]
			p++
		}
		res += makeArrayIncrease(buf[:p], stack)
	}

	return res
}

func makeArrayIncrease(arr []int, stack []int) int {
	// min ops to make arr increas
	// find the longest increasing sub seqence, and the result is n - l
	var p int
	n := len(arr)
	for i := 0; i < n; i++ {
		j := search(p, func(j int) bool {
			return stack[j] > arr[i]
		})
		stack[j] = arr[i]
		if j == p {
			p++
		}
	}

	return n - p
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
