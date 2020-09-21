package p1589

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	var res int
	prev := make([]int, 2)
	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
		cnt := 1 + i/2
		if i&1 == 0 {
			//
			res += cnt*sum - prev[1]
		} else {
			res += cnt*sum - prev[0]
		}
		//res += arr[i]
		if i&1 == 0 {
			prev[0] += sum
		} else {
			prev[1] += sum
		}
	}

	return res
}
