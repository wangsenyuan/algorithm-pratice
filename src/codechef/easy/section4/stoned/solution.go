package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		H := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(H, Q)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}

	fmt.Print(buf.String())
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

const LOG = 20

func solve(H []int, Q [][]int) []int {
	n := len(H)
	P := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		P[i] = make([]int, LOG)
	}

	for j := 0; j < LOG; j++ {
		P[n][j] = n
	}

	D := make([]int, n+1)
	cnt := make([]int, n)
	// node n is root
	D[n] = 0
	stack := make([]int, n)
	var p int
	for i := n - 1; i >= 0; i-- {
		for p > 0 && H[stack[p-1]] <= H[i] {
			p--
		}
		if p == 0 {
			P[i][0] = n
		} else {
			P[i][0] = stack[p-1]
			cnt[i] = cnt[P[i][0]] + 1
		}
		D[i] = D[P[i][0]] + 1
		stack[p] = i
		p++
	}

	for i := 0; i < n; i++ {
		for j := 1; j < LOG; j++ {
			P[i][j] = P[P[i][j-1]][j-1]
		}
	}

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		// D[u] >= D[v]
		for i := LOG - 1; i >= 0; i-- {
			if D[u]-(1<<uint(i)) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := LOG - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	res := make([]int, len(Q))

	for i, cur := range Q {
		u, v := cur[0], cur[1]
		u--
		v--
		if u > v {
			u, v = v, u
		}
		x := lca(u, v)
		res[i] = cnt[x]
		if x != v {
			res[i]++
		}
	}

	return res
}
