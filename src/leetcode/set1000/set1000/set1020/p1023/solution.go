package p1023

func queryString(S string, N int) bool {
	vis := make(map[int]bool)

	for i := 0; i < len(S); i++ {
		if S[i] == '0' {
			continue
		}
		var num int
		for j := i; j < len(S); j++ {
			num = num*2 + int(S[j]-'0')
			if num > 0 && num <= N && !vis[num] {
				vis[num] = true
			}
			if num > N {
				break
			}
		}
	}

	return len(vis) == N
}
