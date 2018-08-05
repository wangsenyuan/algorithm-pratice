package p884

func decodeAtIndex(S string, K int) string {
	n := len(S)
	var curLen int
	for i := 0; i < n; i++ {
		if S[i] >= '0' && S[i] <= '9' {
			d := int(S[i] - '0')
			if curLen*d >= K {
				return decodeAtIndex(S[0:i], (K-1)%curLen+1)
			}
			curLen *= d
		} else {
			curLen++
			if curLen == K {
				return S[i : i+1]
			}
		}
	}

	return ""
}
