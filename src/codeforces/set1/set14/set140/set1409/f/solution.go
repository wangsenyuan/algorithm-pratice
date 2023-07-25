package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	s := readString(reader)[:n]
	t := readString(reader)[:2]
	res := solve(s, t, k)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 28

func solve(s string, t string, k int) int {
	// 如果t[0] = t[1]
	//  只需要让s中出现尽可能多的t[0]即可， 位置不重要
	// 如果t[0] != t[1]
	// dp[i][a][j] 表示到i时，经过j次操作后，有a个t[0]时的计数
	// 如果s[i+1] == t[0]
	// 不变的时候 dp[i+1][a+1][j] = max(dp[i][a][j])
	// 如果s[i+1] == t[1]
	//   dp[i+1][a][j] = max(dp[i][a][j] + a)
	// 或者将其变成t[1]时
	// dp[i+1][a][j+1] = max(dp[i][a][j] + a)
	if t[0] == t[1] {
		return solve1(s, k, t)
	}
	// t[0] != t[1]
	n := len(s)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		// 最多n个a
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := n; j >= 0; j-- {
			for x := k; x >= 0; x-- {
				// 这个必须在第一步
				if s[i] == t[1] {
					dp[j][x] += j
				}
				if s[i] == t[0] && j > 0 {
					// 在保留它为t[0]的时候
					dp[j][x] = max(dp[j][x], dp[j-1][x])
				}
				if x > 0 {
					// 如果将当前s[i]改变成t[1]
					dp[j][x] = max(dp[j][x], dp[j][x-1]+j)
				}

				if j > 0 && x > 0 {
					// 将s[i]变更为t[0]
					dp[j][x] = max(dp[j][x], dp[j-1][x-1])
				}
			}
		}
	}
	var best int
	for a := 0; a <= n; a++ {
		for x := 0; x <= k; x++ {
			best = max(best, dp[a][x])
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}

func solve1(s string, k int, t string) int {
	n := len(s)
	var cnt int
	for i := 0; i < n; i++ {
		if s[i] == t[0] {
			cnt++
		}
	}
	cnt = min(n, cnt+k)

	ans := cnt * (cnt - 1) / 2
	return ans
}
