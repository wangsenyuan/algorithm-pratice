package p1272

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	var res [][]int

	x, y := toBeRemoved[0], toBeRemoved[1]

	var i int

	for i < len(intervals) && intervals[i][1] <= x {
		res = append(res, intervals[i])
		i++
	}

	if i == len(intervals) {
		return res
	}
	// intervals[i][1] > x
	if intervals[i][0] < x {
		res = append(res, []int{intervals[i][0], x})
	}
	for i < len(intervals) && intervals[i][1] <= y {
		i++
	}

	if i == len(intervals) {
		return res
	}
	// intervals[i][1] > y
	if intervals[i][0] < y {
		res = append(res, []int{y, intervals[i][1]})
		i++
	}

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}
