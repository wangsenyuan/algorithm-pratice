package p3443

import "math"

const inf = 1 << 30

func minCostGoodCaption(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([][26]int, n+1)
	minJ := make([]int, n+1)
	nxt := make([][26]int, n+1)
	for i := n - 1; i >= 0; i-- {
		mn := math.MaxInt
		for j := 0; j < 26; j++ {
			res := f[i+1][j] + abs(int(s[i]-'a')-j)
			res2 := math.MaxInt
			if i <= n-6 {
				res2 = f[i+3][minJ[i+3]] + abs(int(s[i]-'a')-j) + abs(int(s[i+1]-'a')-j) + abs(int(s[i+2]-'a')-j)
			}
			if res2 < res || res2 == res && minJ[i+3] < j {
				res = res2
				nxt[i][j] = minJ[i+3] // 记录转移来源
			} else {
				nxt[i][j] = j // 记录转移来源
			}
			f[i][j] = res
			if res < mn {
				mn = res
				minJ[i] = j // 记录最小的 f[i][j] 中的 j 是多少
			}
		}
	}

	ans := make([]byte, n)
	i, j := 0, minJ[0]
	for i < n {
		ans[i] = 'a' + byte(j)
		k := nxt[i][j]
		if k == j {
			i++
		} else {
			ans[i+1] = ans[i]
			ans[i+2] = ans[i]
			i += 3
			j = k
		}
	}
	return string(ans)
}

func abs(num int) int {
	return max(num, -num)
}
