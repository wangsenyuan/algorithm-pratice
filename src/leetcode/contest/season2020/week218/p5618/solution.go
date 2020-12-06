package p5618

func maxOperations(nums []int, k int) int {
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	var res int

	for x, v := range cnt {
		y := k - x

		if x > y {
			continue
		}
		if x == y {
			res += v / 2
			continue
		}
		// x < y
		w := cnt[y]
		res += min(v, w)
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
