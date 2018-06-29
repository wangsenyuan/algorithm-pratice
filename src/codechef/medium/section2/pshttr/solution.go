package main

import (
	"bufio"
	"fmt"
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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 3)
		}
		Q := readNum(scanner)
		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 3)
		}
		res := solve(n, edges, Q, queries)

		for _, ans := range res {
			fmt.Println(ans)
		}

		tc--
	}
}

func solve(n int, edges [][]int, Q int, queries [][]int) []int {
	neighbors := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]int)
	}

	for _, edge := range edges {
		u, v, c := edge[0]-1, edge[1]-1, edge[2]
		neighbors[u][v] = c
		neighbors[v][u] = c
	}

	st := make([]int, n)
	ed := make([]int, n)

	var dfs func(p, u int, time *int)

	dfs = func(p, u int, time *int) {
		(*time)++
		st[u] = *time
		for v := range neighbors[u] {
			if v == p {
				continue
			}
			dfs(u, v, time)
		}
		ed[u] = *time
	}
	dfs(0, 0, new(int))

	ps := make(Pairs, 0, 2*n)
	for i := 0; i < n-1; i++ {
		u, v, c := edges[i][0]-1, edges[i][1]-1, edges[i][2]
		if st[u] < st[v] {
			u, v = v, u
		}
		u++
		// make sure u is larger than 0
		ps = append(ps, Pair{c, -u})
	}

	for i := 0; i < Q; i++ {
		k := queries[i][2]
		ps = append(ps, Pair{k, i})
	}

	res := make([]int, Q)
	sort.Sort(ps)
	bit := CreateBit(n)
	for i := 0; i < len(ps); i++ {
		p := ps[i]
		if p.second < 0 {
			node := -p.second
			node--
			bit.Add(st[node], p.first)
			bit.Add(ed[node]+1, p.first)
		} else {
			idx := p.second
			u, v := queries[idx][0]-1, queries[idx][1]-1
			res[idx] = bit.Get(st[u]) ^ bit.Get(st[v])
		}
	}

	return res
}

type BIT struct {
	arr  []int
	size int
}

func CreateBit(n int) BIT {
	arr := make([]int, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Add(x, v int) {
	for x <= bit.size {
		bit.arr[x] ^= v
		x += x & (-x)
	}
}

func (bit BIT) Get(x int) int {
	var ans int
	for x > 0 {
		ans ^= bit.arr[x]
		x -= x & (-x)
	}
	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

const MAX_LEVEL = 21

func solve1(n int, edges [][]int, Q int, queries [][]int) []int {
	neighbors := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]int)
	}

	for _, edge := range edges {
		u, v, c := edge[0]-1, edge[1]-1, edge[2]
		neighbors[u][v] = c
		neighbors[v][u] = c
	}

	P := make([]int, n)
	PP := make([][]int, n)
	for i := 0; i < n; i++ {
		PP[i] = make([]int, MAX_LEVEL)
	}

	begin := make([]int, n)
	end := make([]int, n)
	// S := make([]int, n)
	var dfs func(p, u int, time *int)

	dfs = func(p, u int, time *int) {
		*time++
		begin[u] = *time

		P[u] = p
		PP[u][0] = p
		for level := 1; level < MAX_LEVEL; level++ {
			t := PP[u][level-1]
			if t > 0 {
				PP[u][level] = PP[t][level-1]
			}
		}

		for v := range neighbors[u] {
			if v == p {
				continue
			}
			dfs(u, v, time)
		}
		end[u] = *time
	}

	dfs(0, 0, new(int))

	es := make(Edges, n-1)

	for i := 0; i < n-1; i++ {
		edge := edges[i]
		u, v, c := edge[0]-1, edge[1]-1, edge[2]
		if P[u] == v {
			u, v = v, u
		}
		// u is parent of v
		es[i] = Edge{u, v, c}
	}

	// sort by cost
	sort.Sort(es)

	isAncestor := func(u, v int) bool {
		return begin[u] <= begin[v] && end[v] <= end[u]
	}

	isInPath := func(p, u, v int) bool {
		return isAncestor(p, u) && isAncestor(u, v)
	}

	lca := func(u, v int) int {
		if isAncestor(u, v) {
			return u
		}

		if isAncestor(v, u) {
			return v
		}

		x := u
		for i := MAX_LEVEL - 1; i >= 0; i-- {
			p := PP[x][i]
			if !isAncestor(p, v) {
				x = p
			}
		}

		return PP[x][0]
	}

	res := make([]int, Q)

	for i := 0; i < Q; i++ {
		query := queries[i]
		u, v, k := query[0]-1, query[1]-1, query[2]
		p := lca(u, v)
		j := sort.Search(len(es), func(j int) bool {
			return es[j].c > k
		})
		var ans int
		for a := 0; a < j; a++ {
			x, y, c := es[a].p, es[a].u, es[a].c
			if isInPath(p, x, u) && isInPath(p, y, u) {
				// p -> x -> y -> u
				ans ^= c
			} else if isInPath(p, x, v) && isInPath(p, y, v) {
				// p -> x -> y -> v
				ans ^= c
			}
			// no other cases
		}
		res[i] = ans
	}

	return res
}

type Edge struct {
	p, u, c int
}

type Edges []Edge

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].c < edges[j].c
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}
