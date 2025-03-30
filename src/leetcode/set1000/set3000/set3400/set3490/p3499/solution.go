package p3499

func maxActiveSectionsAfterTrade(s string) int {
	var sum int
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			sum++
		}
	}

	type pair struct {
		first  int
		second int
	}

	var best int
	var stack []pair
	for i := 0; i < n; i++ {
		if i == 0 || s[i] != s[i-1] {
			stack = append(stack, pair{i, i})
		} else {
			stack[len(stack)-1].second = i
		}
		if s[i] == '0' && len(stack) >= 3 {
			a := stack[len(stack)-3]
			b := stack[len(stack)-1]
			best = max(best, a.second-a.first+1+b.second-b.first+1)
		}
	}

	return sum + best
}
