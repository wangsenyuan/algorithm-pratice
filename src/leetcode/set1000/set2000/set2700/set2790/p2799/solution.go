package p2799

import "strings"

const N = 110

func minimumString(a string, b string, c string) string {
	arr := []string{a, b, c}
	var dfs func(s string, flag int) string

	update := func(ans, tmp string) string {
		if len(ans) == 0 || len(ans) > len(tmp) || len(ans) == len(tmp) && tmp < ans {
			ans = tmp
		}
		return ans
	}

	dfs = func(s string, flag int) string {
		if flag == 7 {
			return s
		}
		ans := ""
		for i := 0; i < 3; i++ {
			if (flag>>i)&1 == 0 {
				var tmp string
				if strings.Contains(s, arr[i]) {
					tmp = dfs(s, flag|(1<<i))
				} else {
					for j := len(arr[i]) - 1; j >= 0; j-- {
						if strings.HasSuffix(s, arr[i][:j+1]) {
							tmp = dfs(s+arr[i][j+1:], flag|(1<<i))
							break
						}
					}
				}
				if tmp == "" {
					tmp = dfs(s+arr[i], flag|(1<<i))
				}
				ans = update(ans, tmp)

			}
		}
		return ans
	}

	return dfs("", 0)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
