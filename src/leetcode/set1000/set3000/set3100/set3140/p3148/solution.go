package p3148

func maxPointsInsideSquare(points [][]int, s string) int {
	// 显然正方形越大，里面的点越多
	n := len(points)

	check := func(dist int) bool {
		var flag int
		for i := 0; i < n; i++ {
			tmp := max(abs(points[i][0]), abs(points[i][1]))
			if tmp > dist {
				continue
			}
			x := int(s[i] - 'a')
			if (flag>>x)&1 == 1 {
				return false
			}
			flag |= 1 << x
		}
		return true
	}

	l, r := 0, 1<<30

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	var res int
	for i := 0; i < n; i++ {
		tmp := max(abs(points[i][0]), abs(points[i][1]))
		if tmp < r {
			res++
		}
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
