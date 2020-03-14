package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(scanner, 2)
		}
		solver := NewSolver(n, A, E)

		var v, k int

		for q > 0 {
			q--
			a, b := readTwoNums(scanner)
			v ^= a
			k ^= b
			v, k = solver.Query(v, k)
			fmt.Printf("%d %d\n", v, k)
		}
	}
}

const MAX_N = 200005

var nodes []*Node
var visit []bool
var G [][]int

func init() {
	nodes = make([]*Node, MAX_N)
	visit = make([]bool, MAX_N)
	G = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		G[i] = make([]int, 0, 3)
	}
}

type Solver struct {
	nodes []*Node
}

func NewSolver(n int, W []int, E [][]int) Solver {
	for i := 0; i < n; i++ {
		visit[i] = false
		G[i] = G[i][:0]
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	var dfs func(v int)

	dfs = func(v int) {
		visit[v] = true
		nodes[v] = NewNode(v, W[v])
		nodes[v].length = 20

		for _, u := range G[v] {
			if !visit[u] {
				dfs(u)
				nodes[v] = merge(nodes[v], nodes[u])
			}
		}
	}

	dfs(0)

	return Solver{nodes[:n]}
}

func (solver Solver) Query(v, k int) (int, int) {
	nodes := solver.nodes

	v--
	node := nodes[v]

	j := 1 << 19

	for node.left != nil || node.right != nil {
		if k&j > 0 && node.left != nil || node.right == nil {
			node = node.left
		} else {
			node = node.right
		}
		j >>= 1
	}
	v = node.vertex + 1
	k ^= node.weight

	return v, k
}

type Node struct {
	left, right *Node
	vertex      int
	weight      int
	length      int
}

func NewNode(vertex, weight int) *Node {
	node := new(Node)
	node.vertex = vertex
	node.weight = weight
	node.length = 0
	return node
}

func (a *Node) Split() {
	tmp := NewNode(a.vertex, a.weight)
	tmp.length = a.length - 1
	if a.weight&(1<<uint(a.length-1)) > 0 {
		a.right = tmp
	} else {
		a.left = tmp
	}
	a.length = 0
	a.vertex = 0
}

func merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.length > 0 {
		a.Split()
	}

	if b.length > 0 {
		b.Split()
	}

	if a.left == nil && a.right == nil {
		if a.vertex > b.vertex {
			return b
		}
		return a
	}

	tmp := new(Node)
	tmp.left = merge(a.left, b.left)
	tmp.right = merge(a.right, b.right)
	tmp.length = 0
	return tmp
}
