package p1238

func circularPermutation(n int, start int) []int {
	N := 1 << uint(n)
	res := make([]int, N)

	var loop func(i int, start, end int)

	loop = func(i int, start, end int) {
		// first half [0.... 011111], [100000, 1111111]
		// then swap second half
		// range for pos i, is start & end
		if i == n {
			// start == end
			return
		}
		mid := (start + end) / 2
		loop(i+1, start, mid)
		loop(i+1, mid, end)
		for j := mid; j < end; j++ {
			res[j] = 1<<uint(n-i-1) | res[j]
		}
		// swap [mid, end)
		for u, v := mid, end-1; u < v; u, v = u+1, v-1 {
			res[u], res[v] = res[v], res[u]
		}
	}

	loop(0, 0, N)

	// find start pos
	var i int

	for i < N && res[i] != start {
		i++
	}

	if i > 0 {
		swap(res, 0, i-1)
		swap(res, i, N-1)
		swap(res, 0, N-1)
	}

	return res
}

func swap(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
