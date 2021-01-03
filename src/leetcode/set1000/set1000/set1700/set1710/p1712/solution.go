package p1712

func minOperations(target []int, arr []int) int {
	m := len(target)
	n := len(arr)

	ii := make(map[int]int)

	for i := 0; i < m; i++ {
		ii[target[i]] = i
	}

	pp := make([]int, 0, m)

	for i := 0; i < n; i++ {
		if j, found := ii[arr[i]]; found {
			pp = append(pp, j)
		}
	}

	stack := make([]int, len(pp))
	var p int

	for i := 0; i < len(pp); i++ {
		j := pp[i]

		k := search(p, func(k int) bool {
			return stack[k] >= j
		})
		if p == k {
			stack[p] = j
			p++
		} else {
			stack[k] = j
		}
	}

	return m - p
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n

	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
