package p3394

import "slices"

func checkValidCuts(n int, rectangles [][]int) bool {
	m := len(rectangles)
	rects := make([]rect, m)
	for i := range m {
		rects[i] = rect{rectangles[i][0], rectangles[i][2], i}
	}
	if check(n, rects) {
		return true
	}

	for i := range m {
		rects[i] = rect{rectangles[i][1], rectangles[i][3], i}
	}
	return check(n, rects)
}

type rect struct {
	l  int
	r  int
	id int
}

func check(inf int, arr []rect) bool {
	slices.SortFunc(arr, func(a, b rect) int {
		return a.l - b.l
	})
	// 在i的前面最小的r是多少
	n := len(arr)
	dp := make([]bool, n)
	r := arr[0].r
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].l < r {
			dp[arr[i].id] = dp[arr[j].id]
			r = max(r, arr[i].r)
			i++
		}
		if i == n {
			break
		}
		// i < n
		dp[arr[i].id] = true
		r = arr[i].r
	}

	slices.SortFunc(arr, func(a, b rect) int {
		return b.r - a.r
	})

	l := arr[0].l

	for i := 0; i < n; {
		for i < n && arr[i].r > l {
			l = min(arr[i].l, l)
			i++
		}
		if i == n {
			return false
		}
		if dp[arr[i].id] {
			return true
		}
		l = arr[i].l
	}

	return false
}
