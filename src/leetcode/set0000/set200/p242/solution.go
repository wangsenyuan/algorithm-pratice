package p242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
		cnt[int(t[i]-'a')]--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
