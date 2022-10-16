package p2436

func countTime(time string) int {

	buf := []byte(time)

	var dfs func(i int) int

	dfs = func(i int) int {
		if i == len(buf) {
			if check(buf) {
				return 1
			}
			return 0
		}
		if buf[i] == '?' {
			var res int
			var end int
			if i == 0 {
				end = 3
			} else if i == 1 {
				end = 10
			} else if i == 3 {
				end = 6
			} else if i == 4 {
				end = 10
			}
			for j := 0; j < end; j++ {
				buf[i] = byte('0' + j)
				res += dfs(i + 1)
			}
			buf[i] = '?'
			return res
		}
		return dfs(i + 1)
	}

	return dfs(0)
}

func check(time []byte) bool {
	hour := int(time[0]-'0')*10 + int(time[1]-'0')
	if hour > 23 {
		return false
	}
	min := int(time[3]-'0')*10 + int(time[4]-'0')
	if min > 59 {
		return false
	}
	return true
}
