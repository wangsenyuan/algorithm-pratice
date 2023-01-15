package p2537

func countGood(nums []int, k int) int64 {
	n := len(nums)
	// n <= 1e5
	// 如果 sub_arr is good, 所有包含它的都是good

	K := int64(k)

	var res int64

	cnt := make(map[int]int)
	var sum int64
	for l, r := 0, 0; r < n; r++ {
		sum -= count(cnt[nums[r]])
		cnt[nums[r]]++
		sum += count(cnt[nums[r]])
		if sum < K {
			continue
		}
		for sum >= K {
			sum -= count(cnt[nums[l]])
			cnt[nums[l]]--
			sum += count(cnt[nums[l]])
			l++
			if sum < K {
				l--
				sum -= count(cnt[nums[l]])
				cnt[nums[l]]++
				sum += count(cnt[nums[l]])
				break
			}
		}
		res += int64(l + 1)
	}

	return res
}

func count(n int) int64 {
	return int64(n) * int64(n-1) / 2
}
