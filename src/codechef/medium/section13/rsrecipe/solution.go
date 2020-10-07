package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	trips := make([][]int, m)
	for i := 0; i < m; i++ {
		trips[i] = readNNums(reader, 3)
	}

	res := solve(n, m, trips)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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

func solve(n, m int, trips [][]int) []int64 {
	degree := make([]int, n+1)

	for i := 0; i < m; i++ {
		t := trips[i]
		x, y := t[0], t[1]
		degree[x-1]++
		degree[y]++
	}

	G := make([][]Edge, n+1)
	//G[0] = make([]Edge, 0, n)
	for i := 0; i <= n; i++ {
		G[i] = make([]Edge, 0, degree[i])
	}

	for i := 0; i < m; i++ {
		t := trips[i]
		x, y, z := t[0], t[1], t[2]
		G[x-1] = append(G[x-1], Edge{x - 1, y, z})
		G[y] = append(G[y], Edge{y, x - 1, -z})
	}

	sums := make([]int64, n+1)
	vis := make([]int, n+1)

	var dfs func(u int, sum int64) bool

	dfs = func(u int, sum int64) bool {
		if vis[u] > 0 {
			return sum == sums[u]
		}
		vis[u]++
		sums[u] = sum

		for _, e := range G[u] {
			if !dfs(e.v, sum+int64(e.cost)) {
				return false
			}
		}
		vis[u]++
		return true
	}

	for i := 0; i <= n; i++ {
		if vis[i] == 0 {
			if !dfs(i, 0) {
				return nil
			}
		}
	}

	res := make([]int64, n)

	for i := 0; i < n; i++ {
		res[i] = sums[i+1] - sums[i]
	}

	return res
}

type Edge struct {
	u, v int
	cost int
}
