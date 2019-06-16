package p898

func subarrayBitwiseORs(A []int) int {
	res := make(map[int]bool)
	cur := make(map[int]bool)
	for i := 0; i < len(A); i++ {
		tmp := make(map[int]bool)
		tmp[A[i]] = true

		for num := range cur {
			tmp[num|A[i]] = true
		}
		cur = tmp
		for num := range cur {
			res[num] = true
		}
	}

	return len(res)
}
