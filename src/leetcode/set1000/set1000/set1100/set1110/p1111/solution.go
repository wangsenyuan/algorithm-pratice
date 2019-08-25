package p1111

func maxDepthAfterSplit(seq string) []int {
	n := len(seq)

	set := make([]int, n)
	var a int
	var b int
	var ma int
	var mb int

	for i := 0; i < n; i++ {
		if seq[i] == '(' {
			if a <= b {
				a++
				set[i] = 0
			} else {
				b++
				set[i] = 1
			}
			if a > ma {
				ma = a
			}
			if b > mb {
				mb = b
			}
		} else {
			if a >= b {
				a--
				set[i] = 0
			} else {
				b--
				set[i] = 1
			}
		}
	}

	return set
}
