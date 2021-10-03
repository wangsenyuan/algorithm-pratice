package p2025

/**
if no change performed, left == right, we can find the answer (0, or x, when mid is 0)
if k == 0
*/
func waysToPartition(nums []int, k int) int {
	cnt := make(map[int64]int)

	var sum int64

	for i := len(nums) - 1; i >= 0; i-- {
		sum += int64(nums[i])
		cnt[sum]++
	}
	var res int
	if sum%2 == 0 {
		res = cnt[sum/2]
		if sum == 0 {
			res--
		}
	}
	cnt2 := make(map[int64]int)
	sum1 := sum
	var sum2 int64

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		cnt[sum1]--
		sum1 -= int64(num)

		// if newSum = sum - num + k, so we need cnt of newSum / 2
		newSum := sum - int64(num) + int64(k)
		if newSum%2 == 0 {
			tmp := cnt2[newSum/2] + cnt[newSum/2]
			if tmp > res {
				res = tmp
			}
		}
		sum2 += int64(num)
		cnt2[sum2]++
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
