package p1964

import "sort"

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	stack := make([]int, n)
	var p int
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		j := sort.Search(p, func(j int) bool {
			return stack[j] > obstacles[i]
		})
		ans[i] = j + 1
		stack[j] = obstacles[i]
		if j == p {
			p++
		}
	}
	return ans
}
