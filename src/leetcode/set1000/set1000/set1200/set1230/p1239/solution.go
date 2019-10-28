package p1239

func maxLength(arr []string) int {
	flags := make([]int, len(arr))
	var n int

	for i := 0; i < len(arr); i++ {
		var flag int
		m := len(arr[i])

		for j := 0; j < m; j++ {
			x := int(arr[i][j] - 'a')
			if flag&(1<<uint(x)) > 0 {
				flag = 0
				break
			}
			flag |= 1 << uint(x)
		}
		if flag == 0 {
			continue
		}
		arr[n] = arr[i]
		flags[n] = flag
		n++
	}
	N := 1 << uint(n)

	var best int

	for state := 1; state < N; state++ {
		var tmp int
		var totalLen int
		var duplicate bool
		for i := 0; i < n; i++ {
			if state&(1<<uint(i)) > 0 {
				flag := flags[i]

				if tmp&flag > 0 {
					duplicate = true
					break
				}
				tmp |= flag
				totalLen += len(arr[i])
			}
		}
		if duplicate {
			continue
		}
		if totalLen > best {
			best = totalLen
		}
	}
	return best
}
