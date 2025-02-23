package p3464

import (
	"slices"
	"sort"
)

func maxDistance(side int, points [][]int, k int) (ans int) {
	type pair struct{ x, y int }
	var a, b []pair
	for _, p := range points {
		x, y := p[0], p[1]
		q := pair{x + y, y - x}
		if x == 0 && y >= 0 || y == side {
			a = append(a, q)
		} else {
			b = append(b, q)
		}
	}

	cmp := func(a, b pair) int { return a.x - b.x }
	slices.SortFunc(a, cmp)
	slices.SortFunc(b, cmp)
	if len(a) == 0 {
		// 保证 a 至少有一个点
		a, b = b, a
	}

	ans = sort.Search(side*2, func(low int) bool {
		low++
		for _, p := range a {
			// 绕一圈
			// 上半圈（从左往右找）
			firstX, firstY := p.x, p.y // 第一个点
			left := k - 1
			curX, lastY := p.x, p.y
			for left > 0 {
				j := sort.Search(len(a), func(i int) bool { return a[i].x >= curX+low })
				if j == len(a) {
					break
				}
				curX = a[j].x
				lastY = a[j].y
				left--
			}
			if left == 0 {
				return false
			}

			// 下半圈最右边的点
			j := sort.Search(len(b), func(i int) bool {
				return max(abs(b[i].x-curX), abs(b[i].y-lastY)) < low
			}) - 1
			// 不能和第一个点离得太近
			if j < 0 || max(abs(b[j].x-firstX), abs(b[j].y-firstY)) < low {
				continue
			}

			// 下半圈（从右往左找）
			left--
			curX = b[j].x
			for left > 0 {
				j := sort.Search(len(b), func(i int) bool { return b[i].x > curX-low }) - 1
				// 不能和第一个点离得太近
				if j < 0 || max(abs(b[j].x-firstX), abs(b[j].y-firstY)) < low {
					break
				}
				curX = b[j].x
				left--
			}
			if left == 0 {
				return false
			}
		}
		return true
	})
	return
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func abs(num int) int {
	return max(num, -num)
}
