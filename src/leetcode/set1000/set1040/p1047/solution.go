package p1047

func removeDuplicates(S string) string {
	n := len(S)
	bs := []byte(S)

	for {
		var j int
		var k int
		for i := 1; i <= n; i++ {
			if i == n || bs[i] != bs[i-1] {
				if (i - k) % 2 == 1 {
					// keep
					bs[j] = bs[i-1]
					j++
				}
				k = i
			}
		}
		if j == n {
			break
		}
		n = j
	}

	return string(bs[:n])
}
