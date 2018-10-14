package p922

func sortArrayByParityII(A []int) []int {
	n := len(A)
	B := make([]int, n)
	// ps := make(Pairs, n)
	var j int
	for i := 0; i < n; i++ {
		num := A[i]
		if num&1 == 1 {
			B[j] = num
			j++
		}
	}
	for i := 0; i < n; i++ {
		num := A[i]
		if num&1 == 0 {
			B[j] = num
			j++
		}
	}

	// [odd ... even]
	a := 0
	b := n / 2
	if b&1 == 0 {
		b++
	}
	for b < n {
		B[a], B[b] = B[b], B[a]
		a += 2
		b += 2
	}

	return B
}

// type Pair struct {
// 	first, second int
// }

// type Pairs []Pair

// func (ps Pairs) Len() int {
// 	return len(ps)
// }

// func (ps Pairs) Less(i, j int) bool {
// 	return ps[i].first < ps[j].first
// }

// func (ps Pairs) Swap(i, j int) {
// 	ps[i], ps[j] = ps[j], ps[i]
// }
