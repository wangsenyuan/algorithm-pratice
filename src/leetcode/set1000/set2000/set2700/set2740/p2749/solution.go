package p2749

const MOD = 1e9 + 7

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func numberOfGoodSubarraySplits(nums []int) int {
	res := 1
	prev := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if prev >= 0 {
			res = mul(res, i-prev)
		}
		prev = i
	}
	if prev < 0 {
		return 0
	}
	return res
}
