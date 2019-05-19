package p1047

func removeDuplicates(S string) string {
	n := len(S)
	bs := []byte(S)

	var j int

	for i, k := 1, 0; i <= n; i++ {
		if i == n || bs[i] != bs[i-1] {
			if (i-k)%2 == 1 {
				if j > 0 && bs[j-1] == bs[i-1] {
					// also need to remove
					j--
				} else {
					// keep it
					bs[j] = bs[i-1]
					j++
				}
			}
			k = i
		}
	}

	return string(bs[:j])
}

func removeDuplicates1(S string) string {
	n := len(S)
	bs := []byte(S)

	for {
		var j int
		var k int
		for i := 1; i <= n; i++ {
			if i == n || bs[i] != bs[i-1] {
				if (i-k)%2 == 1 {
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
