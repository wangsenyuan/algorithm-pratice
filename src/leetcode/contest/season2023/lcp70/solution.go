package lcp70

func sandyLandManagement(size int) [][]int {
	var res [][]int

	for i := 1; i < size; i++ {
		for j := 1; j <= 2*i-1; j += 4 {
			res = append(res, []int{i, j})
		}
	}
	res = append(res, []int{size, 1})

	if size > 1 {
		res = append(res, []int{size, size*2 - 1})
	}

	for _, j := range []int{4, 7, 9} {
		for i := j; i < 2*size-1; i += 8 {
			res = append(res, []int{size, i})
		}
	}

	return res
}
