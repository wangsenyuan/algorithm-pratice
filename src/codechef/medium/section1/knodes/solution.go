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

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)

		edges := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}

		K := readNum(scanner)
		queries := make([][]int, K)

		for i := 0; i < K; i++ {
			m := readNum(scanner)
			queries[i] = readNNums(scanner, m)
		}

		res := solve(n, edges, K, queries)

		for _, ans := range res {
			if ans {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}

		t--
	}
}

func solve(n int, edges [][]int, K int, queries [][]int) []bool {
	graph := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		graph[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		graph[u][v] = true
		graph[v][u] = true
	}

	maxLevel := 21
	parent := make([][]int, n)
	for i := 0; i < n; i++ {
		parent[i] = make([]int, maxLevel)
	}

	begin := make([]int, n)
	end := make([]int, n)

	var dfs func(p, u int, time *int)

	dfs = func(p, u int, time *int) {
		(*time)++
		begin[u] = *time
		parent[u][0] = p

		for level := 1; level < maxLevel; level++ {
			pp := parent[u][level-1]
			if pp == 0 {
				parent[u][level] = 0
				continue
			}
			parent[u][level] = parent[pp][level-1]
		}

		for v := range graph[u] {
			if v == p {
				continue
			}
			dfs(u, v, time)
		}
		end[u] = *time
	}

	dfs(0, 0, new(int))

	isAncestor := func(u, v int) bool {
		return begin[u] <= begin[v] && end[v] <= end[u]
	}

	lca := func(u, v int) int {
		if isAncestor(u, v) {
			return u
		}

		if isAncestor(v, u) {
			return v
		}
		node := u
		for i := maxLevel - 1; i >= 0; i-- {
			p := parent[node][i]
			if p > 0 && !isAncestor(p, v) {
				node = p
			}
		}
		return parent[node][0]
	}

	// u ---> v ----> w
	liesInPath := func(u, v, w int) bool {
		return isAncestor(u, v) && isAncestor(v, w)
	}

	res := make([]bool, K)

	for i := 0; i < K; i++ {
		query := queries[i]
		lnode := query[0] - 1
		rnode := 0
		for j := 0; j < len(query); j++ {
			node := query[j] - 1
			if begin[rnode] < begin[node] {
				rnode = node
			}

			if isAncestor(lnode, node) || isAncestor(node, lnode) {
				// in the same path
				if isAncestor(lnode, node) {
					lnode = node
				}
			} else if begin[lnode] > begin[node] {
				lnode = node
			}
		}

		pnode := lca(lnode, rnode)
		yes := true

		for i := 0; i < len(query); i++ {
			node := query[i] - 1
			if lnode == rnode {
				// lnode is bottom one
				yes = yes && isAncestor(node, lnode)
			} else {
				yes = yes && (liesInPath(pnode, node, lnode) || liesInPath(pnode, node, rnode))
			}
		}
		res[i] = yes
	}

	return res
}

func solve1(n int, edges [][]int, K int, queries [][]int) []bool {
	graph := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		graph[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		graph[u][v] = true
		graph[v][u] = true
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	level := make([]int, n)

	que := make([]int, n)
	front, end := 0, 0
	que[end] = 0
	end++

	for front < end {
		u := que[front]
		front++

		for v := range graph[u] {
			if v == 0 || parent[v] != -1 {
				// already visited
				continue
			}
			parent[v] = u
			level[v] = level[u] + 1
			que[end] = v
			end++
		}
	}

	nodes := make(Nodes, n)

	for i := 0; i < n; i++ {
		nodes[i] = Node{i, level[i]}
	}

	sort.Sort(nodes)

	flag := make([]bool, n)

	// how many branchs end at i
	branchs := make([]int, n)

	check := func() bool {
		for i := 0; i < n; i++ {
			branchs[i] = 0
		}
		// from bottom to top
		for _, node := range nodes {
			u := node.label
			for v := range graph[u] {
				if flag[u] && branchs[v] == 2 {
					// 3 branchs at v
					return false
				}
				branchs[u] += branchs[v]
			}
			if branchs[u] > 2 {
				// 3 branchs below u
				return false
			}
			if branchs[u] == 0 && flag[u] {
				// a new branch from u
				branchs[u] = 1
			}
		}
		return true
	}

	res := make([]bool, K)
	for k := 0; k < K; k++ {
		nds := queries[k]
		for i := 0; i < n; i++ {
			flag[i] = false
		}

		for _, nd := range nds {
			flag[nd-1] = true
		}

		res[k] = check()
	}

	return res
}

type Node struct {
	label int
	level int
}

type Nodes []Node

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
	// high at front
	return nodes[i].level > nodes[j].level
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}
