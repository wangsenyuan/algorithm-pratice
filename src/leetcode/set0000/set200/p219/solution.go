package p219

func containsNearbyDuplicate(nums []int, k int) bool {
	idx := make(map[int]int)

	for i, num := range nums {
		if j, has := idx[num]; has && i-j <= k {
			return true
		}
		idx[num] = i
	}
	return false
}
