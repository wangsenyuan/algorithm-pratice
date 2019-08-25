package p1096

import (
	"sort"
)

func braceExpansionII(expression string) []string {
	n := len(expression)
	dp := make([]int, n)
	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		if expression[i] == '{' {
			stack[p] = i
			p++
		} else if expression[i] == '}' {
			dp[stack[p-1]] = i
			p--
		} else if expression[i] >= 'a' && expression[i] <= 'z' {
			dp[i] = i
		}
	}

	union := func(a, b []string) []string {
		set := make(map[string]bool)
		var res []string
		for _, x := range a {
			set[x] = true
			res = append(res, x)
		}

		for _, y := range b {
			if !set[y] {
				res = append(res, y)
			}
		}
		return res
	}

	concat := func(a, b []string) []string {
		var res []string
		for _, x := range a {
			for _, y := range b {
				res = append(res, x+y)
			}
		}
		return res
	}

	var loop func(start, end int) []string

	loop = func(start, end int) []string {
		if start == end {
			return []string{expression[start : end+1]}
		}
		var un bool
		if dp[start] == end {
			start++
			end--
			un = true
		}
		i := start
		var res []string
		for i <= end {
			j := dp[i]
			fn := concat
			if un {
				if expression[i] == ',' {
					i++
				}
				j = dp[i] + 1
				fn = union
				for j <= end && expression[j] != ',' {
					j = dp[j] + 1
				}
				j--
			}

			tmp := loop(i, j)
			if len(res) == 0 {
				res = tmp
			} else {
				res = fn(res, tmp)
			}
			i = j + 1
		}

		return res
	}

	res := loop(0, n-1)
	sort.Strings(res)

	return res
}
