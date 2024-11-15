package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	s := fmt.Sprintf("%v", res)

	fmt.Print(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 3)
	}
	qs := readNNums(reader, m)
	return solve(n, edges, qs)
}

func solve(n int, edges [][]int, queries []int) []int {
	slices.SortFunc(edges, func(a, b []int) int {
		return a[2] - b[2]
	})

	type pair struct {
		first  int
		second int
	}
	m := len(queries)
	qs := make([]pair, m)
	for i, q := range queries {
		qs[i] = pair{q, i}
	}

	slices.SortFunc(qs, func(a, b pair) int {
		return a.first - b.first
	})

	set := NewDSU(n)

	count := func(sz int) int {
		return (sz - 1) * sz / 2
	}

	ans := make([]int, m)

	var sum int

	for i, j := 0, 0; i < m; i++ {
		for j < len(edges) && edges[j][2] <= qs[i].first {
			edge := edges[j]
			u, v := edge[0]-1, edge[1]-1
			u = set.Find(u)
			v = set.Find(v)
			// 它们肯定没有连起
			sum -= count(set.cnt[u])
			sum -= count(set.cnt[v])
			set.Union(u, v)
			u = set.Find(u)
			sum += count(set.cnt[u])
			j++
		}
		ans[qs[i].second] = sum
	}

	return ans
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
