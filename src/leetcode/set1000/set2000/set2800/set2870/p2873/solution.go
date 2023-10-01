package p2873

const H = 22

func maxSubarrays(nums []int) int {
	n := len(nums)

	pref := make([]int, n+1)
	pref[0] = (1 << H) - 1
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] & nums[i]
	}

	if pref[n] != 0 {
		return 1
	}
	suf := (1 << H) - 1
	var ans int
	for i := n - 1; i >= 0; i-- {
		suf &= nums[i]
		if suf == 0 && (pref[i] == 0 || i == 0) {
			ans++
			suf = (1 << H) - 1
		}
	}
	return ans
}
