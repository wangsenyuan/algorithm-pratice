package p128

func longestConsecutive(nums []int) int {
	m := make(map[int]int)

	res := 0
	for _, num := range nums {
		if _, has := m[num]; has {
			continue
		}

		left, right := m[num-1], m[num+1]
		m[num] = left + right + 1
		m[num-left] = left + right + 1
		m[num+right] = left + right + 1

		if left+right+1 > res {
			res = left + right + 1
		}
	}

	return res
}
