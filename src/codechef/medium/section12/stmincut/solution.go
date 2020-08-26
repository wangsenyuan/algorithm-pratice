package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}
		fmt.Println(solve(n, A))
	}
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

const H = 20

func solve(n int, A [][]int) int64 {

	var res int64
	ps := make([]Pair2, 0, n*(n+1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := max(A[i][j], A[j][i])
			b := min(A[i][j], A[j][i])
			res += int64(a - b)
			A[i][j] = a
			A[j][i] = a
			ps = append(ps, Pair2{A[i][j], Pair{i, j}})
		}
	}
	sort.Sort(Pair2Array(ps))

	ds := NewDS(n)

	adj := make([][]Pair, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0, n)
	}

	for _, p := range ps {
		w := p.first
		u := p.second.first
		v := p.second.second

		if ds.Union(u, v) {
			adj[u] = append(adj[u], Pair{v, w})
			adj[v] = append(adj[v], Pair{u, w})
		}
	}

	level := make([]int, n)
	pp := make([][]int, n)
	val := make([][]int, n)

	for i := 0; i < n; i++ {
		pp[i] = make([]int, H)
		val[i] = make([]int, H)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		for i := 1; i < H; i++ {
			pp[u][i] = pp[pp[u][i-1]][i-1]
			val[u][i] = min(val[u][i-1], val[pp[u][i-1]][i-1])
		}

		for _, x := range adj[u] {
			v := x.first
			w := x.second
			if v == p {
				continue
			}

			level[v] = level[u] + 1
			pp[v][0] = u
			val[v][0] = w
			dfs(u, v)
		}
	}

	dfs(-1, 0)

	ff := func(u, v int) int {
		if level[u] < level[v] {
			u, v = v, u
		}
		var res = 1 << 30
		for i := H - 1; i >= 0; i-- {
			if level[u]-1<<uint(i) >= level[v] {
				res = min(res, val[u][i])
				u = pp[u][i]
			}
		}
		if u == v {
			return res
		}

		for i := H - 1; i >= 0; i-- {
			if pp[u][i] != pp[v][i] {
				res = min(res, val[u][i])
				res = min(res, val[v][i])
				u = pp[u][i]
				v = pp[v][i]
			}
		}

		res = min(res, val[u][0])
		res = min(res, val[v][0])

		return res
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res += int64(ff(i, j)-A[i][j]) * 2
		}
	}

	return res
}

type DS struct {
	arr []int
	cnt []int
}

func NewDS(n int) DS {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return DS{arr, cnt}
}

func (ds *DS) Find(x int) int {
	if ds.arr[x] != x {
		ds.arr[x] = ds.Find(ds.arr[x])
	}
	return ds.arr[x]
}

func (ds *DS) Union(x, y int) bool {
	px := ds.Find(x)
	py := ds.Find(y)

	if px == py {
		return false
	}
	if ds.cnt[px] < ds.cnt[py] {
		px, py = py, px
	}

	ds.arr[py] = px
	ds.cnt[px] += ds.cnt[py]
	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Pair2 struct {
	first  int
	second Pair
}

type Pair2Array []Pair2

func (arr Pair2Array) Len() int {
	return len(arr)
}

func (arr Pair2Array) Less(i, j int) bool {
	return arr[i].first > arr[j].first
}

func (arr Pair2Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
