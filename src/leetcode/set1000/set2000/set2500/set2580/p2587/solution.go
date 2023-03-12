package p2587

func beautifulSubarrays(nums []int) int64 {
	var res int64

	cnt := make(map[int]int)

	var x int
	cnt[0]++

	for _, num := range nums {
		x ^= num
		res += int64(cnt[x])
		cnt[x]++
	}

	return res
}
