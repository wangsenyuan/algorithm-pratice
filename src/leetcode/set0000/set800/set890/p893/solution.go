package p893

import "sort"

func numSpecialEquivGroups(A []string) int {
	if len(A) == 0 {
		return 0
	}
	m := len(A[0])
	s1 := make([]byte, m/2)
	s2 := make([]byte, m-m/2)
	mem := make(map[string]int)

	for i := 0; i < len(A); i++ {
		for j := 0; j < m; j++ {
			if j%2 == 0 {
				s2[j/2] = A[i][j]
			} else {
				s1[j/2] = A[i][j]
			}
		}
		sort.Sort(ByteSlice(s1))
		sort.Sort(ByteSlice(s2))
		a := string(s1)
		b := string(s2)
		c := a + b
		mem[c]++
	}

	return len(mem)
}

type ByteSlice []byte

func (bs ByteSlice) Len() int {
	return len(bs)
}

func (bs ByteSlice) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs ByteSlice) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
