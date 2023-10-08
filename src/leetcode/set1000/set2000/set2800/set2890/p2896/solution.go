package p2896

import "sort"

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)
	sort.Ints(tasks)

	reverse(processorTime)

	n := len(processorTime)

	var ans int
	for i := 0; i < n; i++ {
		for j := 4 * i; j < 4*(i+1); j++ {
			ans = max(ans, processorTime[i]+tasks[j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
