package p3483

func totalNumbers(digits []int) int {
	cnt := make([]int, 10)
	for _, x := range digits {
		cnt[x]++
	}
	var res int

	var m int
	for _, v := range cnt {
		if v > 0 {
			m++
		}
	}

	// n := len(digits)
	for i := 1; i < 10; i++ {
		if cnt[i] > 0 {
			// 这个作为第一位
			for j := 0; j < 10; j += 2 {
				if cnt[j] == 0 || i == j && cnt[i] == 1 {
					continue
				}
				// 这个作为最后一位
				if i != j {
					// 中间的选其他的
					res += m - 2
					// 中间选择i
					if cnt[i] > 1 {
						res++
					}
					// 中间选择j
					if cnt[j] > 1 {
						res++
					}
				} else {
					// i == j
					// 选择其他的
					res += m - 1
					if cnt[i] >= 3 {
						res++
					}
				}
			}
		}
	}

	return res
}
