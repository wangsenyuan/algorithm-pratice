package p3191

func minOperations(nums []int) int {
	n := len(nums)

	cnt := make([]int, n+3)

	var res int

	for i, num := range nums {
		if i > 0 {
			cnt[i] += cnt[i-1]
		}
		if cnt[i]%2 == 1 {
			num ^= 1
		}

		if num == 0 {
			if i+3 > n {
				return -1
			}
			res++
			cnt[i]++
			cnt[i+3]--
		}
	}

	return res
}
