package p2716

func minimizedStringLength(s string) int {
	cnt := make([]int, 26)
	var res int

	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		cnt[x]++
		if cnt[x] == 1 {
			res++
		}
	}
	return res
}
