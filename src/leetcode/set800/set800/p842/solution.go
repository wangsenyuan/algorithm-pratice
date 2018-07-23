package p842

import (
	"strconv"
)

const INF = 1 << 31

func splitIntoFibonacci(S string) []int {

	var dfs func(a, b int, S string) bool

	var res []int

	dfs = func(a, b int, S string) bool {
		if len(S) == 0 {
			return len(res) > 0
		}

		c := a + b
		if c < 0 || c >= INF {
			return false
		}
		cs := strconv.Itoa(c)
		var i int
		for i < len(S) && i < len(cs) && S[i] == cs[i] {
			i++
		}

		if i < len(cs) {
			return false
		}
		res = append(res, c)
		return dfs(b, c, S[i:])
	}
	n := len(S)
	var a int
	for i := 1; i < n; i++ {
		if S[0] == '0' && i > 1 {
			break
		}
		a = a*10 + int(S[i-1]-'0')
		if a < 0 || a >= INF {
			break
		}
		var b int
		for j := i + 1; j < n; j++ {
			if S[i] == '0' && j-i > 2 {
				break
			}
			b = b*10 + int(S[j-1]-'0')
			if b < 0 || b >= INF {
				break
			}
			res = make([]int, 0, 20)

			if dfs(a, b, S[j:]) {
				res = append([]int{a, b}, res...)
				return res
			}
		}
	}

	return make([]int, 0)
}
