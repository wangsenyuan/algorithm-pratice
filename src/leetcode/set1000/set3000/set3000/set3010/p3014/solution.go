package p3014

func countOfPairs(n int, x int, y int) []int64 {
	if x > y {
		x, y = y, x
	}

	diff := make([]int, n+2)
	add := func(l, r, v int) {
		if l > r {
			return
		}
		diff[l] += v
		diff[r+1] -= v
	}

	update := func(i, x, y int) {
		add(y-i, n-i, -1) // 撤销 [y,n]
		dec := y - x - 1  // 缩短的距离
		add(y-i-dec, n-i-dec, 1)

		j := (x+y+1)/2 + 1
		add(j-i, y-1-i, -1) // 撤销 [j, y-1]
		add(x-i+2, x-i+y-j+1, 1)
	}

	update2 := func(i, x, y int) {
		add(y-i, n-i, -1)            // 撤销 [y,n]
		dec := (y - i) - (i - x + 1) // 缩短的距离
		add(y-i-dec, n-i-dec, 1)

		j := i + (y-x+1)/2 + 1
		add(j-i, y-1-i, -1) // 撤销 [j, y-1]
		add(i-x+2, i-x+y-j+1, 1)
	}

	for i := 1; i <= n; i++ {
		add(1, i-1, 1)
		add(1, n-i, 1)
		if x+1 >= y {
			continue
		}
		if i <= x {
			update(i, x, y)
		} else if i >= y {
			update(n+1-i, n+1-y, n+1-x)
		} else if i < (x+y)/2 {
			update2(i, x, y)
		} else if i > (x+y+1)/2 {
			update2(n+1-i, n+1-y, n+1-x)
		}
	}

	ans := make([]int64, n)
	sumD := int64(0)
	for i, d := range diff[1 : n+1] {
		sumD += int64(d)
		ans[i] = sumD
	}
	return ans
}
