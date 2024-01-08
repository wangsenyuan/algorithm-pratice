package p3000

func areaOfMaxDiagonal(dimensions [][]int) int {
	best := []int{0, 0}

	for _, cur := range dimensions {
		h, w := cur[0], cur[1]
		if h*h+w*w > best[0] || h*h+w*w == best[0] && h*w > best[1] {
			best = []int{h*h + w*w, h * w}
		}
	}
	return best[1]
}
