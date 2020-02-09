package p1343

func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	sum := make([]int64, n+1)
	K := int64(k)
	T := int64(threshold)

	var res int

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(arr[i])

		if i+1 >= k {
			aa := sum[i+1] - sum[i+1-k]
			if T*K <= aa {
				res++
			}
		}
	}

	return res
}
