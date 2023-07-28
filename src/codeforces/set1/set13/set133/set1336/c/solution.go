package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	x := readString(reader)
	res := solve(s, x)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const mod = 998244353

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(s string, t string) int {
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][i] = 1
	}

	for ln := 1; ln <= n; ln++ {
		c := s[ln-1]
		for i := 0; i+ln <= n; i++ {
			j := i + ln
			if i >= m || t[i] == c {
				dp[i][j] = add(dp[i][j], dp[i+1][j])
			}
			if j > m || t[j-1] == c {
				dp[i][j] = add(dp[i][j], dp[i][j-1])
			}
		}
	}
	var ans int
	for i := m; i <= n; i++ {
		ans = add(ans, dp[0][i])
	}
	return ans
}

func solve1(s string, t string) int {
	n := len(s)
	m := len(t)
	// m <= n
	/**
	  dp[i][j] = 处理到s[i]时， 从目标字符长a开始，满足a是t的一个子串且从位置j开始时的方案数
	  a的长度 = min(i, m - j) 这个是确定的
	*/

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[n][0] = 1

	for i := n - 1; i >= 1; i-- {
		for j := 0; j <= m; j++ {
			if j == 0 {
				if i >= m {
					dp[i][0] = n - i + 1
				} else if s[i] == t[i] {
					dp[i][0] = dp[i+1][0]
				}
			} else if j == m {
				dp[i][m] = mul(2, dp[i+1][m])
				if s[i] == t[m-1] {
					dp[i][m] = add(dp[i][m], dp[i+1][m-1])
				}
			} else {
				if i+j >= m || s[i] == t[i+j] {
					dp[i][j] = dp[i+1][j]
				}
				if s[i] == t[j-1] {
					dp[i][j] = add(dp[i][j], dp[i+1][j-1])
				}
			}
		}
	}
	ans := dp[1][m]
	for i := 0; i < m; i++ {
		if t[i] == s[0] {
			ans = add(ans, dp[1][i])
		}
	}
	ans = mul(ans, 2)
	return ans
}
