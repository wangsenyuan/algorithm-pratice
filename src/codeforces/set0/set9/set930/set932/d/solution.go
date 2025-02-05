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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	queries := make([][]int, n)
	for i := range n {
		queries[i] = readNNums(reader, 3)
	}
	return solve(queries)
}

const inf = 1e18

func solve(queries [][]int) []int {

	// node[0] = inf
	nodes := []int{inf, 0}

	// 最多n个node
	n := len(queries)

	H := bits.Len(uint(n))

	fa := make([][]int, n+1)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		fa[i] = make([]int, H)
		dp[i] = make([]int, H)
	}
	// dp[0][?] = inf
	for j := 0; j < H; j++ {
		dp[1][j] = inf
	}

	var ans []int
	var last int

	findParent := func(u int, w int) int {
		// 保证 nodes[u] < w 始终成立
		for i := H - 1; i >= 0; i-- {
			if nodes[fa[u][i]] < w {
				u = fa[u][i]
			}
		}
		return fa[u][0]
	}

	add := func(p int, q int) {
		r := last ^ p
		w := last ^ q

		m := len(nodes)

		nodes = append(nodes, w)
		if nodes[r] < w && r > 0 {
			r = findParent(r, w)
		}

		fa[m][0] = r

		dp[m][0] = nodes[r]

		for i := 1; i < H; i++ {
			fa[m][i] = fa[fa[m][i-1]][i-1]
			// avoid overflow
			dp[m][i] = min(inf, dp[m][i-1]+dp[fa[m][i-1]][i-1])
		}
	}

	get := func(p int, q int) int {
		u := last ^ p
		x := last ^ q
		if nodes[u] > x {
			// dp[u][0] = nodes[u]
			return 0
		}
		x -= nodes[u]
		cnt := 1
		for i := H - 1; i >= 0; i-- {
			if dp[u][i] <= x {
				cnt += 1 << i
				x -= dp[u][i]
				u = fa[u][i]
			}
		}

		return cnt
	}

	for _, cur := range queries {
		if cur[0] == 1 {
			add(cur[1], cur[2])
		} else {
			last = get(cur[1], cur[2])
			ans = append(ans, last)
		}
	}

	return ans
}
