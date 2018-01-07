package p760

func anagramMappings(A []int, B []int) []int {
	idx := make(map[int][]int)

	for i := 0; i < len(B); i++ {
		if is, found := idx[B[i]]; found {
			idx[B[i]] = append(is, i)
		} else {
			idx[B[i]] = make([]int, 0, 10)
			idx[B[i]] = append(idx[B[i]], i)
		}
	}
	res := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		js := idx[A[i]]
		res[i] = js[len(js)-1]
		idx[A[i]] = js[:len(js)-1]
	}

	return res
}
