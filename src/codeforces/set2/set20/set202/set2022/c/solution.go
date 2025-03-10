package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	buf.WriteTo(os.Stdout)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) int {
	readString(reader)
	a := make([]string, 2)
	a[0] = readString(reader)
	a[1] = readString(reader)
	return solve(len(a[0]), a)
}

const inf = 1 << 60

func solve(n int, a []string) int {
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, 5)
		for j := range 5 {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	get := func(x byte, y byte, z byte) int {
		var cnt int
		if x == 'A' {
			cnt++
		} else {
			cnt--
		}
		if y == 'A' {
			cnt++
		} else {
			cnt--
		}
		if z == 'A' {
			cnt++
		} else {
			cnt--
		}
		// 不会等于0
		if cnt > 0 {
			return 1
		}
		return 0
	}

	for i := range n {
		// state 表示的是i前面两列的状态
		for state := range 5 {
			if state == 0 {
				// 前面都填好了，这里有3种选择
				if i+3 <= n {
					// 上下两排分别是两个选区
					// i, i + 1, i + 2
					dp[i+3][0] = max(dp[i+3][0], dp[i][0]+get(a[0][i], a[0][i+1], a[0][i+2])+get(a[1][i], a[1][i+1], a[1][i+2]))
				}
				if i+2 <= n {
					// 或者放置一个L型
					dp[i+2][1] = max(dp[i+2][1], dp[i][0]+get(a[0][i], a[1][i], a[1][i+1]))
					// 或者放置一个倒L型
					dp[i+2][3] = max(dp[i+2][3], dp[i][0]+get(a[0][i], a[1][i], a[0][i+1]))
				}
			} else if state == 1 {
				// 有两个选择
				if i+1 <= n && i > 0 {
					// 放置一个反/倒L, 第i列处理好了, i+1列还没有处理
					dp[i+1][0] = max(dp[i+1][0], dp[i][1]+get(a[0][i-1], a[0][i], a[1][i]))
				}
				if i+2 <= n && i > 0 {
					dp[i+2][4] = max(dp[i+2][4], dp[i][1]+get(a[0][i-1], a[0][i], a[0][i+1]))
				}
			} else if state == 2 {
				// 表示第一行有，前两列都是空的，只有一个选择
				if i >= 2 && i+1 <= n {
					dp[i+1][3] = max(dp[i+1][3], dp[i][2]+get(a[0][i-2], a[0][i-1], a[0][i]))
				}
			} else if state == 3 {
				// 表示前一列第二行还没有匹配，有两个选择
				if i+1 <= n && i > 0 {
					// 放置一个反L
					dp[i+1][0] = max(dp[i+1][0], dp[i][3]+get(a[1][i-1], a[0][i], a[1][i]))
				}
				// 或者第二列放置3个
				if i+2 <= n && i > 0 {
					dp[i+2][2] = max(dp[i+2][2], dp[i][3]+get(a[1][i-1], a[1][i], a[1][i+1]))
				}
			} else if state == 4 {
				// 表示第二行，前两个还没有匹配，只有一种选择
				if i >= 2 && i+1 <= n {
					dp[i+1][1] = max(dp[i+1][1], dp[i][4]+get(a[1][i-2], a[1][i-1], a[1][i]))
				}
			}
		}
	}

	return dp[n][0]
}
