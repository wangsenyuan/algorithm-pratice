package p1442

func countTriplets(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	var res int
	for i := 0; i < n; i++ {
		x := arr[i]
		for j := i + 1; j < n; j++ {
			x ^= arr[j]
			if x != 0 {
				continue
			}
			res += j - i
		}
	}

	return res
}
