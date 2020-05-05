package visitorder

func visitOrder(points [][]int, direction string) []int {
	n := len(points)

	res := make([]int, 0, n)

	var start int

	for i := 1; i < n; i++ {
		if points[start][0] > points[i][0] {
			start = i
		}
	}

	res = append(res, start)

	used := make([]bool, n)

	used[start] = true

	for i := 0; i < len(direction); i++ {
		next := -1

		if direction[i] == 'L' {
			for j := 0; j < n; j++ {
				if used[j] {
					continue
				}
				if next < 0 || orientation(points[start], points[next], points[j]) > 0 {
					next = j
				}
			}
		} else {
			for j := 0; j < n; j++ {
				if used[j] {
					continue
				}
				if next < 0 || orientation(points[start], points[next], points[j]) < 0 {
					next = j
				}
			}
		}
		used[next] = true
		res = append(res, next)
		start = next
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			res = append(res, i)
		}
	}

	return res
}

type Point []int

func orientation(p, q, r Point) int {
	v := (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
	if v == 0 {
		return 0
	}
	if v > 0 {
		// clock direction
		return 1
	}
	// anti-clock
	return -1
}
