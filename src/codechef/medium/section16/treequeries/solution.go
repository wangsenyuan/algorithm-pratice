package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	// tc := 1

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		q := readNum(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}

		res := solve(n, E, Q)

		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const H = 19

func solve(n int, E [][]int, Q [][]int) []int {
	adj := make([][]int, n+1)

	addEdge := func(u, v int) {
		if adj[u] == nil {
			adj[u] = make([]int, 0, 1)
		}
		adj[u] = append(adj[u], v)
	}

	for _, e := range E {
		u, v := e[0], e[1]
		addEdge(u, v)
		addEdge(v, u)
	}

	for i := 1; i <= n; i++ {
		sort.Ints(adj[i])
	}

	tin := make([]int, n+1)
	sz := make([]int, n+1)
	enter := make([]int, n+1)
	anc := make([][]int, n+1)
	left := make([][]int, n+1)
	right := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		anc[i] = make([]int, H)
		left[i] = make([]int, H)
		right[i] = make([]int, H)
		sz[i]++
	}

	var ord []int
	var time int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		tin[u] = time
		time++

		ord = append(ord, u)

		for i := 1; i < H; i++ {
			anc[u][i] = anc[anc[u][i-1]][i-1]
		}

		for _, v := range adj[u] {
			if p == v {
				left[u][0] = time - tin[u]
				continue
			}
			anc[v][0] = u
			dfs(u, v)
			sz[u] += sz[v]
		}

		right[u][0] = left[u][0] + n - sz[u] - 1

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			enter[v] = tin[v] - tin[u]
			if v > p {
				enter[v] += n - sz[u]
			}
		}
	}
	

	dfs(0, 1)

	kth := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if k >= (1 << i) {
				k -= 1 << i
				u = anc[u][i]
			}
		}
		return u
	}

	update := func(table [][]int, u int, i int) {
		table[u][i] += left[u][i-1] + table[anc[u][i-1]][i-1]
		if anc[u][i-1] > 1 {
			x := kth(u, (1<<(i-1))-1)
			if x < anc[x][1] {
				table[u][i] -= sz[x]
			}
		}
	}

	query := func(u int, x int) int {
		for i := H - 1; i >= 0; i-- {
			if anc[u][i] > 0 && x >= left[u][i] && x <= right[u][i] {
				y := kth(u, (1<<i)-1)
				x -= left[u][i]
				if enter[y] <= x {
					x += sz[y]
				}
				u = anc[u][i]
			}
		}

		j := x + tin[u]

		if x >= left[u][0] {
			j -= n - sz[u]
		}

		return ord[j]
	}

	for i := 1; i < H; i++ {
		for j := 1; j <= n; j++ {
			update(left, j, i)
			update(right, j, i)
		}
	}
	var last int
	ans := make([]int, len(Q))

	for i, cur := range Q {
		c, d := cur[0], cur[1]
		a, x := c^last, d^last
		ans[i] = query(a, x)
		last = ans[i]
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
