package p836

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x1, y1, x2, y2 := rec1[0], rec1[1], rec1[2], rec1[3]
	x3, y3, x4, y4 := rec2[0], rec2[1], rec2[2], rec2[3]

	return x3 < x2 && y3 < y2 && x4 > x1 && y4 > y1
}
