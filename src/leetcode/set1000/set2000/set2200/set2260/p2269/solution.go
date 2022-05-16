package p2269

func divisorSubstrings(num int, k int) int {

	origin := num

	var res int

	var arr []int

	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}

	var cur int

	n := len(arr)

	var base int64 = 1

	for i := 0; i < k-1; i++ {
		base *= 10
	}

	for i := n - 1; i >= 0; i-- {
		cur = cur*10 + arr[i]

		if n-i >= k {
			if cur > 0 && origin%cur == 0 {
				res++
			}
			cur = int(int64(cur) % base)
		}
	}

	return res
}
