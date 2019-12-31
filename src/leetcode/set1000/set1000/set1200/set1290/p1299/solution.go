package p1299

func replaceElements(arr []int) []int {
	n := len(arr)
	cur := -1

	for i := n - 1; i >= 0; i-- {
		x := arr[i]
		arr[i] = cur
		cur = max(cur, x)
	}
	return arr
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
