package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, cnt, res := process(reader)
	fmt.Println(cnt)
	fmt.Println(res)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (k int, s string, ops int, res string) {
	s = readString(reader)
	k = readNum(reader)
	ops, res = solve(s, k)
	return
}
func solve(s string, k int) (int, string) {
	n := len(s)
	if k >= n {
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteByte(s[i])
			if i+1 < n {
				buf.WriteByte('+')
			}
		}
		return 0, buf.String()
	}

	// k < n
	// fp[i][j] 表示要将i...j变成回文需要的操作数
	fp := make([][]int, n)
	for i := range n {
		fp[i] = make([]int, n)
	}
	for i := range n {
		for j := i - 1; j >= 0; j-- {
			if j+1 < i {
				fp[j][i] = fp[j+1][i-1]
			}
			if s[j] != s[i] {
				fp[j][i]++
			}
		}
	}
	//dp[j][i] 表示在前缀i上得到最多j个回文时的最小修改值
	dp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = n
		}
	}

	for i := range n {
		dp[1][i] = fp[0][i]
	}

	for j := 2; j <= k; j++ {
		for i := range n {
			for u := i; u > 0; u-- {
				dp[j][i] = min(dp[j][i], dp[j-1][u-1]+fp[u][i])
			}
			// 如果可以用更少的分段，得到更好的结果时，更优
			dp[j][i] = min(dp[j][i], dp[j-1][i])
		}
	}

	ops := dp[k][n-1]

	var dfs func(i int, j int) string

	dfs = func(i int, j int) string {
		w := dp[j][i]
		for j > 1 && dp[j-1][i] == w {
			j--
		}
		// 这么多个分段
		// j > 1
		u := 0
		if j > 1 {
			u = 1
			for dp[j-1][u-1]+fp[u][i] != w {
				u++
			}
		}
		buf := make([]byte, i-u+1)
		for l, r := u, i; l <= r; l, r = l+1, r-1 {
			buf[l-u] = s[l]
			buf[r-u] = s[r]
			if buf[l-u] != buf[r-u] {
				buf[r-u] = buf[l-u]
			}
		}
		if j == 1 {
			return string(buf)
		}
		return dfs(u-1, j-1) + "+" + string(buf)
	}

	res := dfs(n-1, k)

	return ops, res
}
