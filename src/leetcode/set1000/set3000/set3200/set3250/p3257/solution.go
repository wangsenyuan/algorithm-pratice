package p3257

func countKConstraintSubstrings(s string, k int) int {

	n := len(s)
	var res int
	cnt := make([]int, 2)
	for l, r := 0, 0; r < n; r++ {
		cnt[int(s[r]-'0')]++
		for cnt[0] > k && cnt[1] > k {
			cnt[int(s[l]-'0')]--
			l++
		}
		res += r + 1 - l
	}
	return res
}
