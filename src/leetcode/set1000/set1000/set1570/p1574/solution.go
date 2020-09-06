package p1574

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	cur := j
	// j is the starting point of non-decreasing suffix
	var best = j
	for i := 0; i < cur; i++ {
		// if we delete after i, the answer is
		for j < n && arr[j] < arr[i] {
			j++
		}
		if j-i-1 < best {
			best = j - i - 1
		}
		if arr[i] > arr[i+1] {
			break
		}
	}
	return best
}
