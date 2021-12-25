package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		firstLine := readNNums(scanner, 3)
		H, W, Q := firstLine[0], firstLine[1], firstLine[2]
		grid := make([][]int, H)
		for i := 0; i < H; i++ {
			grid[i] = readNNums(scanner, W)
		}
		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 3)
		}
		res := solve(H, W, grid, Q, queries)

		for i := 0; i < Q; i++ {
			fmt.Println(res[i])
		}
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(H, W int, grid [][]int, Q int, queries [][]int) []int {
	qs := make([]Query, Q)

	for i := 0; i < Q; i++ {
		qs[i] = Query{queries[i][0] - 1, queries[i][1] - 1, queries[i][2], i}
	}

	sort.Sort(Queries(qs))

	cells := make([]*Cell, H*W)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			cells[i*W+j] = &Cell{i, j, grid[i][j]}
		}
	}

	sort.Sort(Cells(cells))

	N := H * W
	uf := NewUF(N)

	vis := make([]bool, H*W)

	ans := make([]int, Q)

	for i, j := 0, 0; i < Q; i++ {
		query := qs[i]

		for j < N && cells[j].p < query.p {
			x, y := cells[j].x, cells[j].y

			for k := 0; k < 4; k++ {
				u, v := x+dd[k], y+dd[k+1]
				if u >= 0 && u < H && v >= 0 && v < W {
					if vis[u*W+v] {
						uf.Union(x*W+y, u*W+v)
					}
				}
			}

			vis[x*W+y] = true
			j++
		}

		if !vis[query.r*W+query.c] {
			ans[query.i] = 0
			continue
		}
		pp := uf.Find(query.r*W + query.c)

		ans[query.i] = uf.Count(pp)
	}

	return ans
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(size int) UF {
	arr := make([]int, size)
	cnt := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UF{arr, cnt}
}

func (uf UF) Count(x int) int {
	return uf.cnt[x]
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] >= uf.cnt[pb] {
		uf.arr[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	} else {
		uf.arr[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	}
	return true
}

type Cell struct {
	x, y, p int
}

type Cells []*Cell

func (cells Cells) Len() int {
	return len(cells)
}

func (cells Cells) Less(i, j int) bool {
	return cells[i].p < cells[j].p
}

func (cells Cells) Swap(i, j int) {
	cells[i], cells[j] = cells[j], cells[i]
}

type Query struct {
	r int
	c int
	p int
	i int
}

type Queries []Query

func (this Queries) Len() int {
	return len(this)
}

func (this Queries) Less(i, j int) bool {
	return this[i].p < this[j].p
}

func (this Queries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
