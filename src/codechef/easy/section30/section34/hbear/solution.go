package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	G := make([]string, n)
	for i := 0; i < n; i++ {
		G[i] = readString(reader)
	}
	q := readNum(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, m, k, G, Q)

	var buf bytes.Buffer
	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n, m, k int, G []string, Q [][]int) []int {
	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == 'H' {
				sum[i][j]++
			}
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	// dp[i][j][a] = the largest b, ending at position i, j, having width a
	dp := make([][]int, m+5)
	for i := 0; i <= m+4; i++ {
		dp[i] = make([]int, n+5)
	}

	check := func(x1, y1, x2, y2 int) bool {
		tmp := sum[x2][y2]
		if x1 > 0 {
			tmp -= sum[x1-1][y2]
		}
		if y1 > 0 {
			tmp -= sum[x2][y1-1]
		}
		if x1 > 0 && y1 > 0 {
			tmp += sum[x1-1][y1-1]
		}
		return tmp >= k
	}

	add := func(x1, y1, x2, y2 int) {
		if y1 > y2 || x1 > x2 {
			return
		}
		dp[x1][y1]++
		dp[x2+1][y1]--
		dp[x1][y2+1]--
		dp[x2+1][y2+1]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p := m
			lst := m - 1
			for k := i; k < n; k++ {
				for p-1 >= j && check(i, j, k, p-1) {
					p--
				}
				if p > lst {
					continue
				}
				add(k-i+1, p-j+1, n-i, lst-j+1)
				lst = p - 1
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	ans := make([]int, len(Q))
	for i, cur := range Q {
		a, b := cur[0], cur[1]
		ans[i] = dp[a][b]
	}
	return ans
}
