package p1815

const MOD = 1000000007

func countNicePairs(nums []int) int {
	cnt := make(map[int]int)
	var res int
	for i := 0; i < len(nums); i++ {
		tmp := nums[i] - rev(nums[i])
		res += cnt[tmp]
		res %= MOD
		cnt[tmp]++
	}

	return res
}

func rev(num int) int {
	var res int
	for num > 0 {
		r := num % 10
		res = res*10 + r
		num /= 10
	}
	return res
}
