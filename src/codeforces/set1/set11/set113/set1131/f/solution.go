package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	pairs := make([][]int, n-1)

	for i := range n - 1 {
		pairs[i] = readNNums(reader, 2)
	}

	res := solve(n, pairs)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func solve(n int, pairs [][]int) []int {

	sz := make([]int, n*2)

	for i := 0; i < n; i++ {
		sz[i] = 1
	}

	fa := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		fa[i] = i
	}

	var find func(u int) int
	find = func(u int) int {
		if fa[u] != u {
			fa[u] = find(fa[u])
		}
		return fa[u]
	}

	id := n

	g := make([][]int, 2*n)

	for _, cur := range pairs {
		x, y := cur[0]-1, cur[1]-1
		x = find(x)
		y = find(y)
		g[id] = append(g[id], x, y)
		fa[x] = id
		fa[y] = id
		sz[id] += sz[x]
		sz[id] += sz[y]
		id++
	}

	res := make([]int, n)

	var dfs func(u int, l int, r int)
	dfs = func(u int, l int, r int) {
		if l == r {
			// u must be leaf
			res[l] = u + 1
			return
		}
		// len(g[u]) = 2
		a, b := g[u][0], g[u][1]

		dfs(a, l, l+sz[a]-1)
		dfs(b, l+sz[a], r)
	}

	dfs(id-1, 0, n-1)

	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
