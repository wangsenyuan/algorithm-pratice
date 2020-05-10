package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	F := make([][]int, m)

	for i := 0; i < m; i++ {
		F[i] = readNNums(reader, 2)
	}

	cnt, res := solve(n, m, F)

	if cnt == 0 {
		fmt.Println(res)
	} else {
		fmt.Println(cnt)
		fmt.Println(res)
	}
}

func solve(n, m int, F [][]int) (int, string) {
	// if F has loop, there there is no answer
	vis := make([]int, n+1)

	outs := make([][]int, n+1)
	ins := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		outs[i] = make([]int, 0, 10)
		ins[i] = make([]int, 0, 10)
	}

	for _, f := range F {
		u, v := f[0], f[1]
		outs[u] = append(outs[u], v)
		ins[v] = append(ins[v], u)
	}

	var dfs func(u int) bool

	dfs = func(u int) bool {
		if vis[u] == 1 {
			// a loop
			return false
		}

		if vis[u] > 0 {
			return true
		}

		vis[u] = 1

		for _, v := range outs[u] {
			if !dfs(v) {
				return false
			}
		}

		vis[u] = 2

		return true
	}

	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			if !dfs(i) {
				return 0, "-1"
			}
		}
	}

	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			return 0, "-1"
		}
	}

	seq := make([]int, n)

	var topo func(u int)

	order := n - 1

	topo = func(u int) {

		for _, v := range outs[u] {
			if vis[v] == 0 {
				topo(v)
			}
		}

		seq[order] = u
		vis[u] = order + 1
		order--
	}

	for i := 1; i <= n; i++ {
		vis[i] = 0
	}

	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			topo(i)
		}
	}

	dp := make([]int, n+1)
	fp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		fp[i] = i
	}

	for _, u := range seq {
		for _, v := range outs[u] {
			dp[v] = min(dp[v], dp[u])
		}
	}

	for i := n - 1; i >= 0; i-- {
		u := seq[i]
		for _, v := range ins[u] {
			fp[v] = min(fp[v], fp[u])
		}
	}

	res := make([]byte, n)
	var cnt int
	for i := 1; i <= n; i++ {
		j := min(dp[i], fp[i])
		if i == j {
			res[i-1] = 'A'
			cnt++
		} else {
			res[i-1] = 'E'
		}
	}

	return cnt, string(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
