package p3468

const inf = 1 << 60

func countArrays(original []int, bounds [][]int) int {
	// c[i] - c[i-1] = o[i] - o[i-1]
	// c[i] - o[i] = c[i-1] - o[i-1]
	// 也就是说幅度要一样，且在允许的幅度访问内
	// c[i] <= b[i][1], c[i] >= b[i][0]
	// d[i] = c[i] - o[i] => d[i] + o[i] <= b[i][1] => d[i] <= b[i][1] - o[i]
	// d[i] + o[i] >= b[i][0] => d[i] >= b[i][0] - o[i]
	n := len(original)
	d := []int{inf, -inf}
	for i := range n {
		d[0] = min(d[0], bounds[i][1]-original[i])
		d[1] = max(d[1], bounds[i][0]-original[i])
	}

	return max(0, d[0]-d[1]+1)
}
