package p171

func titleToNumber(s string) int {
	n := len(s)
	var ans int

	for i := 0; i < n; i++ {
		ans = ans*26 + int(s[i]-'A') + 1
	}
	return ans
}
