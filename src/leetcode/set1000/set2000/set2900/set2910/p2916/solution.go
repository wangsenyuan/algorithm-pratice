package p2916

const H = 31

func findKOr(nums []int, k int) int {
	cnt := make([]int, H)
	for _, num := range nums {
		for i := 0; i < H; i++ {
			cnt[i] += (num >> i) & 1
		}
	}

	var res int
	for i := 0; i < H; i++ {
		if cnt[i] >= k {
			res |= 1 << i
		}
	}

	return res
}
