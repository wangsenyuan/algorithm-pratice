package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, d := readTwoNums(reader)
	X := make([][]int, n)

	for i := 0; i < n; i++ {
		X[i] = readNNums(reader, d)
	}

	fmt.Println(solve(n, d, X))
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

const INF = 1 << 30
const N_INF = -INF

func solve(n, d int, X [][]int) int64 {
	D := 1 << uint(d)

	tmp := make([]Pair, n)

	edges := make([]Edge, 0, 2*D*n)

	addEdge := func(u, v int) {
		if u == v {
			return
		}
		var dist int
		for i := 0; i < d; i++ {
			dist += abs(X[u][i] - X[v][i])
		}
		edges = append(edges, Edge{dist, Pair{u, v}})
	}

	for mask := 0; mask < D; mask++ {
		var p int
		for i := 0; i < n; i++ {
			var sum int
			for j := 0; j < d; j++ {
				if mask&(1<<uint(j)) > 0 {
					sum += X[i][j]
				} else {
					sum -= X[i][j]
				}
			}
			tmp[p] = Pair{sum, i}
			p++
		}

		var first = Pair{INF, -1}
		var second = Pair{N_INF, -1}

		for i := 0; i < n; i++ {
			first = first.min(tmp[i])
			second = second.max(tmp[i])
		}
		for i := 0; i < n; i++ {
			addEdge(i, first.second)
			addEdge(i, second.second)
		}
	}

	sort.Sort(Edges(edges))

	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	var get func(x int) int
	get = func(x int) int {
		if arr[x] != x {
			arr[x] = get(arr[x])
		}
		return arr[x]
	}

	union := func(x int, y int) bool {
		px := get(x)
		py := get(y)
		if px == py {
			return false
		}
		if cnt[px] < cnt[py] {
			px, py = py, px
		}
		cnt[px] += cnt[py]
		arr[py] = px
		return true
	}

	var res int64

	for i := 0; i < len(edges); i++ {
		cur := edges[i]
		if union(cur.vs.first, cur.vs.second) {
			res += int64(cur.dist)
			if cnt[get(cur.vs.first)] == n {
				break
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pair struct {
	first, second int
}

func (this Pair) min(that Pair) Pair {
	if this.first < that.first {
		return this
	}
	return that
}

func (this Pair) max(that Pair) Pair {
	if this.first > that.first {
		return this
	}
	return that
}

type Edge struct {
	dist int
	vs   Pair
}

type Edges []Edge

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].dist > edges[j].dist
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}
