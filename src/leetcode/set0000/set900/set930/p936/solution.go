package p936

func movesToStamp(stamp string, target string) []int {
	m := len(stamp)
	n := len(target)

	str := []byte(target)
	cnt := n

	check := func(i int) bool {
		var hit bool
		for j := 0; j < m; j++ {
			if str[i+j] == '?' {
				continue
			}
			hit = true
			if str[i+j] != stamp[j] {
				return false
			}
		}
		return hit
	}
	res := make([]int, 0, n)
	for cnt > 0 {
		var dec int
		for i := 0; i+m <= n; i++ {
			if check(i) {
				res = append(res, i)
				for j := 0; j < m; j++ {
					if str[i+j] != '?' {
						dec++
						str[i+j] = '?'
					}
				}
				break
			}
		}
		if dec == 0 {
			//find none
			return []int{}
		}

		cnt -= dec
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
