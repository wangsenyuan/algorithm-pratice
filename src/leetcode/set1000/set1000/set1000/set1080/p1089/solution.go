package p1089

func duplicateZeros(arr []int) {
	n := len(arr)

	res := make([]int, 2*n)

	for i, j := 0, 0; i < n; i++ {
		if arr[i] == 0 {
			res[j] = 0
			j++
			res[j] = 0
			j++
		} else {
			res[j] = res[i]
			j++
		}
	}
	copy(arr, res)
}
