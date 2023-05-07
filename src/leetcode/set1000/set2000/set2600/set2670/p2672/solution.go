package p2672

func colorTheArray(n int, queries [][]int) []int {
	// n <= 1e5
	// color <= 1e5
	arr := make([]int, n)
	var count int

	process := func(i int, c int) {
		// change i to color c
		if arr[i] == 0 {
			// first time have color
			if i > 0 && arr[i-1] == c {
				count++
			}
			if i+1 < n && arr[i+1] == c {
				count++
			}
			arr[i] = c
			return
		}
		// already have a color
		if i > 0 && arr[i-1] == arr[i] {
			count--
		}
		if i+1 < n && arr[i+1] == arr[i] {
			count--
		}
		arr[i] = c

		// first time have color
		if i > 0 && arr[i-1] == c {
			count++
		}
		if i+1 < n && arr[i+1] == c {
			count++
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		process(cur[0], cur[1])
		ans[i] = count
	}
	return ans
}
