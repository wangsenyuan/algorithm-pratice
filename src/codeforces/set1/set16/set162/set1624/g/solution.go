package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readString(reader)
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, m, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, m int, E [][]int) int {
	g := NewGraph(n, 2*m)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	vis := make([]int, n)

	var dfs func(u int, mask int)

	dfs = func(u int, mask int) {
		vis[u] = mask
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if (mask&w) != w || vis[v] == mask {
				continue
			}
			dfs(v, mask)
		}
	}

	mask := (1 << 30) - 1

	for d := 29; d >= 0; d-- {
		// if we turn off d, whether connect of not
		tmp := mask ^ (1 << d)

		dfs(0, tmp)

		ok := true

		for i := 0; i < n; i++ {
			if vis[i] != tmp {
				ok = false
				break
			}
		}

		if ok {
			mask = tmp
		}
	}

	return mask
}

func solve1(n int, m int, E [][]int) int {
	g := NewGraph(n, 2*m)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	uf := NewUF(n)

	var dfs func(p int, u int, mask int)

	dfs = func(p int, u int, mask int) {

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if p == v || (mask&w) != w || uf.Find(v) == uf.Find(u) {
				continue
			}
			uf.Union(u, v)
			dfs(u, v, mask)
		}
	}

	mask := (1 << 30) - 1

	for d := 29; d >= 0; d-- {
		// if we turn off d, whether connect of not
		tmp := mask ^ (1 << d)

		uf.Reset()
		dfs(-1, 0, tmp)

		if uf.Size() == 1 {
			// good
			mask = tmp
		}
	}

	return mask
}

type UF struct {
	arr  []int
	cnt  []int
	size int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &UF{arr, cnt, n}
}

func (this *UF) Reset() {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] = i
		this.cnt[i] = 1
	}
	this.size = len(this.arr)
}

func (this *UF) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *UF) Union(a, b int) bool {
	pa := this.Find(a)
	pb := this.Find(b)

	if pa == pb {
		return false
	}
	if this.cnt[pa] < this.cnt[pb] {
		pa, pb = pb, pa
	}
	this.cnt[pa] += this.cnt[pb]
	this.arr[pb] = pa
	this.size--
	return true
}

func (this *UF) Size() int {
	return this.size
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
