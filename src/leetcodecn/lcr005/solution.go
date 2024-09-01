package lcr005

func maxProduct(words []string) int {
	type pair struct {
		first  int
		second int
	}
	n := len(words)
	vals := make([]pair, n)

	for i, w := range words {
		var flag int
		for j := 0; j < len(w); j++ {
			flag |= 1 << int(w[j]-'a')
		}
		vals[i] = pair{flag, len(w)}
	}
	var res int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if vals[i].first&vals[j].first == 0 {
				res = max(res, vals[i].second*vals[j].second)
			}
		}
	}
	return res
}
