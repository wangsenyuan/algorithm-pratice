package p2306

func distinctNames(ideas []string) int64 {
	occ := make(map[string]bool)
	for _, idea := range ideas {
		occ[idea] = true
	}
	// leading x cnt so far
	cnt := make([]int, 26)
	cnt2 := make([][]int, 26)
	for i := 0; i < 26; i++ {
		cnt2[i] = make([]int, 26)
	}
	var res int64
	n := len(ideas)

	buf := make([]byte, 10)

	for i := 0; i < n; i++ {
		idea := ideas[i]
		x := int(idea[0] - 'a')
		for j := 0; j < 26; j++ {
			if x == j {
				continue
			}
			buf[0] = byte(j + 'a')
			copy(buf[1:], []byte(idea[1:]))
			if occ[string(buf[:len(idea)])] {
				// 如果把j交互过来，在ideas中存在，则不符合条件
				continue
			}
			res += int64(cnt[j] - cnt2[j][x])
		}
		for j := 0; j < 26; j++ {
			if x == j {
				continue
			}
			buf[0] = byte(j + 'a')
			copy(buf[1:], []byte(idea[1:]))
			if occ[string(buf[:len(idea)])] {
				cnt2[x][j]++
			}
		}
		cnt[x]++
	}

	return 2 * res
}
