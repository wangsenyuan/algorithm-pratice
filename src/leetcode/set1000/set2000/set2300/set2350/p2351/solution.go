package p2351

func repeatedCharacter(s string) byte {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		cnt[x]++
		if cnt[x] == 2 {
			return s[i]
		}
	}
	return '0'
}
