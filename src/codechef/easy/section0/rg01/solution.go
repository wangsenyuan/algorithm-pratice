package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readOne(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	bs := scanner.Bytes()
	readInt(bs, 0, &a)
	return
}

func readTwo(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	readInt(bs, x+1, &b)
	return
}

func readThree(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	readInt(bs, x+1, &c)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	v, e := readTwo(scanner)

	edges := make([][]int, e)
	for i := 0; i < e; i++ {
		u, v, w := readThree(scanner)
		edges[i] = []int{u, v, w}
	}
	res := solve(v, e, edges)
	for i := 0; i < v; i++ {
		fmt.Printf("%d", res[i][0])
		for j := 1; j < v; j++ {
			fmt.Printf(" %d", res[i][j])
		}
		fmt.Println()
	}
}

func solve(V int, E int, edges [][]int) [][]int {
	sort.Sort(sort.Reverse(Edges(edges)))

	set := make([]int, V)
	rank := make([]int, V)
	for i := 0; i < V; i++ {
		set[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		px := set[x]
		if px != x {
			set[x] = find(px)
		}
		return set[x]
	}

	union := func(x, y int) bool {
		px := find(x)
		py := find(y)
		if px == py {
			return false
		}
		if rank[px] > rank[py] {
			set[py] = px
		} else if rank[px] < rank[py] {
			set[px] = py
		} else {
			set[py] = px
			rank[px]++
		}
		return true
	}

	conns := make([][][]int, V)
	for i := 0; i < V; i++ {
		conns[i] = make([][]int, 0, V)
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if union(u, v) {
			conns[u] = append(conns[u], []int{v, w})
			conns[v] = append(conns[v], []int{u, w})
		}
	}

	res := make([][]int, V)

	for i := 0; i < V; i++ {
		res[i] = make([]int, V)
		for j := 0; j < V; j++ {
			res[i][j] = -1
		}
	}

	que := make([]int, V)
	vis := make([]bool, V)
	bfs := func(x int) {
		for i := 0; i < V; i++ {
			vis[i] = false
		}
		head, tail := 0, 0
		que[tail] = x
		res[x][x] = math.MaxInt32
		vis[x] = true
		tail++
		for head < tail {
			tmp := tail
			for head < tmp {
				v := que[head]
				head++
				for _, conn := range conns[v] {
					u, w := conn[0], conn[1]
					if !vis[u] {
						res[x][u] = min(res[x][v], w)
						vis[u] = true
						que[tail] = u
						tail++
					}
				}
			}
		}
		// res[x][x] = 0
	}

	for i := 0; i < V; i++ {
		bfs(i)
	}

	for i := 0; i < V; i++ {
		res[i][i] = 0
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Edges [][]int

func (this Edges) Len() int {
	return len(this)
}

func (this Edges) Less(i, j int) bool {
	a, b := this[i], this[j]
	return a[2] < b[2]
}

func (this Edges) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
