package p1310

func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	sum := make([]int, n+1)

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] ^ arr[i]
	}

	res := make([]int, len(queries))

	for i, query := range queries {
		l, r := query[0], query[1]
		res[i] = sum[r+1] ^ sum[l]
	}

	return res
}
