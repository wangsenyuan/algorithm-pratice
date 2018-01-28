package p771

func numJewelsInStones(J string, S string) int {
	mp := make(map[byte]bool)

	for i := 0; i < len(J); i++ {
		mp[J[i]] = true
	}

	var res int

	for i := 0; i < len(S); i++ {
		if mp[S[i]] {
			res++
		}
	}
	return res
}
