package p1727

func countGoodRectangles(rectangles [][]int) int {
	var mxLen int

	for _, rec := range rectangles {
		tmp := min(rec[0], rec[1])
		mxLen = max(mxLen, tmp)
	}

	var res int

	for _, rec := range rectangles {
		tmp := min(rec[0], rec[1])
		if tmp == mxLen {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
