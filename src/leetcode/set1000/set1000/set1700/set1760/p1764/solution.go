package p1764

func canChoose(groups [][]int, nums []int) bool {

	for i, j := 0, 0; j < len(groups); j++ {
		g := groups[j]
		var ok bool
		for i+len(g) <= len(nums) {
			var k int
			for k < len(g) && g[k] == nums[i+k] {
				k++
			}
			if k == len(g) {
				i += k
				ok = true
				break
			}
			i++
		}
		if !ok {
			return false
		}
	}

	return true
}
