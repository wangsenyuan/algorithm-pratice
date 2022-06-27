package p2317

const H = 30

func maximumXOR(nums []int) int {
	var res int

	// a ^ (a & b)
	// 1 ^ (1 & 0) = 1
	// 1 ^ (1 & 1) = 0
	// 0 ^ (0 & 1) = 0
	// 0 ^ (0 & 0) = 0

	n := len(nums)

	cnt := make([]int, 2)

	for h := H - 1; h >= 0; h-- {
		// h位是否可以为1,
		for j := 0; j < 2; j++ {
			cnt[j] = 0
		}

		for i := 0; i < n; i++ {
			cnt[(nums[i]>>h)&1]++
		}
		if cnt[1] > 0 {
			res |= 1 << h
		}
	}

	return res
}
