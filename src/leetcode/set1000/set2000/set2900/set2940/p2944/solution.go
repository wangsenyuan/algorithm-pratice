package p2944

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	// 找出最大的连续长度

	var h int
	for i := 0; i < len(hBars); {
		lo := hBars[i]
		i++
		for i < len(hBars) && hBars[i] == hBars[i-1]+1 {
			i++
		}
		hi := hBars[i-1]
		h = max(h, hi-lo+2)
	}

	var w int
	for i := 0; i < len(vBars); {
		l := vBars[i]
		i++
		for i < len(vBars) && vBars[i] == vBars[i-1]+1 {
			i++
		}
		r := vBars[i-1]
		w = max(w, r-l+2)
	}

	sz := min(h, w)
	return sz * sz
}
