package p1779

func nearestValidPoint(x int, y int, points [][]int) int {
	dist := 1 << 30
	pivot := -1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			var tmp int
			if p[0] == x {
				tmp = abs(p[1] - y)
			} else {
				tmp = abs(p[0] - x)
			}
			if dist > tmp {
				pivot = i
				dist = tmp
			}
		}
	}
	return pivot
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
