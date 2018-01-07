package p761

import "sort"

func makeLargestSpecial(S string) string {
	var count int
	ss := make([]string, 0, 10)
	for i, j := 0, 0; i < len(S); i++ {
		if S[i] == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			tmp := makeLargestSpecial(S[j+1 : i])
			ss = append(ss, "1"+tmp+"0")
			j = i + 1
		}
	}
	sort.Strings(ss)
	var res string
	for i := len(ss) - 1; i >= 0; i-- {
		res += ss[i]
	}
	return res
}
