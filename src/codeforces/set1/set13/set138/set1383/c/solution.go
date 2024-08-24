package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		readNum(reader)
		a := readString(reader)
		b := readString(reader)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(a, b string) int {
	g := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		g[i] = make([]bool, 26)
	}

	n := len(a)

	adj := make([]int, 26)

	for i := 0; i < n; i++ {

		// 假设有m棵树，那么答案 = n - m
		// 所以 a[i] = b[i]不应该连起来，否则有可能会减少树的数量
		if a[i] != b[i] {
			// 这棵树里面，每条边就是一次变化
			g[int(a[i]-'a')][int(b[i]-'a')] = true
			g[int(b[i]-'a')][int(a[i]-'a')] = true
			adj[int(a[i]-'a')] |= 1 << int(b[i]-'a')
		}
	}

	vis := make([]bool, 26)

	var dfs func(u int) int
	dfs = func(u int) int {
		vis[u] = true
		res := 1
		for v := 0; v < 26; v++ {
			if g[u][v] && !vis[v] {
				res += dfs(v)
			}
		}
		return res
	}

	var comp int

	for i := 0; i < 20; i++ {
		if !vis[i] {
			dfs(i)
			comp++
		}
	}

	dp := make([]bool, 1<<20)

	dp[0] = true

	var ans int
	for state := 0; state < 1<<20; state++ {
		if dp[state] {
			ans = max(ans, bits.OnesCount(uint(state)))
			for u := 0; u < 20; u++ {
				if (state>>u)&1 == 0 && adj[u]&state == 0 {
					dp[state|(1<<u)] = true
				}
			}
		}
	}

	return 2*20 - comp - ans
}
