package p2555

func substringXorQueries(s string, queries [][]int) [][]int {
	qs := make(map[int][]int)

	for i, cur := range queries {
		x := cur[0] ^ cur[1]
		if len(qs[x]) == 0 {
			qs[x] = make([]int, 0, 1)
		}
		qs[x] = append(qs[x], i)
	}
	ans := make([][]int, len(queries))
	for i := 0; i < len(ans); i++ {
		ans[i] = []int{-1, -1}
	}
	n := len(s)
	for l := 0; l < n; l++ {
		if s[l] == '0' {
			if js, ok := qs[0]; ok {
				for _, j := range js {
					ans[j] = []int{l, l}
				}
				delete(qs, 0)
			}
			continue
		}
		var num int
		for i := 0; i < 30 && l+i < n; i++ {
			num = (num << 1) | int(s[l+i]-'0')
			if js, ok := qs[num]; ok {
				for _, j := range js {
					ans[j] = []int{l, l + i}
				}
			}
			delete(qs, num)
		}
	}

	return ans
}
