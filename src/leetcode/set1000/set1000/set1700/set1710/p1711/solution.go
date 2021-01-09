package p1711

const MOD = 1000000007

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func waysToSplit(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	pref := make([]int, n+1)
	var sum int
	for i := 0; i < n; i++ {
		sum += nums[i]
		pref[i+1] = sum
	}

	var res int
	for i := 2; i < n; i++ {
		//if i is the last index
		right := sum - pref[i]
		j := search(i, func(j int) bool {
			return pref[i]-pref[j] < pref[j]
		}) - 1
		// left := pref[j]
		mid := pref[i] - pref[j]
		if mid <= right {
			k := search(i, func(k int) bool {
				return pref[i]-pref[k] <= right
			}) - 1
			// pref[i] - pref[k] > right
			if k < 0 {
				k = 0
			}

			modAdd(&res, j-k)
		}
	}

	return res
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
