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
		n, m := readTwoNums(scanner)

		edges := make([][]int, m)

		for i := 0; i < m; i++ {
			edges[i] = readNNums(scanner, 2)
		}

		res := solve(n, m, edges)

		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}

		t--
	}
}

func solve(n, m int, edges [][]int) []int {
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = new(Node)
		nodes[i].label = i
		nodes[i].neighbors = make(map[int]bool)
	}

	for i := 0; i < m; i++ {
		u, v := edges[i][0]-1, edges[i][1]-1
		nodes[u].AddNeighbor(v)
		nodes[v].AddNeighbor(u)
	}

	// sort by degree desc
	sort.Sort(Nodes(nodes))

	res := make([]int, n)
	// remove all edges, need to add back n - 1 edges to connect them all
	res[n-1] = n - 1

	uf := CreateUF(n)
	set := make([]bool, n)
	var j int
	for d := n - 1; d > 0; d-- {
		// degree go down from n - 1 to 0
		for j < n && nodes[j].Degree() == d {
			// add them into the set
			u := nodes[j].label
			set[u] = true

			for v := range nodes[j].neighbors {
				if set[v] {
					// v is already in set
					uf.Union(u, v)
				}
			}

			j++
		}
		// need to add edges to connect dis-connected groups
		res[d-1] = uf.Size() - 1
	}

	return res
}

type Node struct {
	label     int
	neighbors map[int]bool
}

func (node *Node) AddNeighbor(x int) {
	node.neighbors[x] = true
}

func (node Node) Degree() int {
	return len(node.neighbors)
}

type Nodes []*Node

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
	return len(nodes[i].neighbors) > len(nodes[j].neighbors)
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

type UF struct {
	set   []int
	count []int
	size  int
}

func CreateUF(n int) UF {
	set := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		count[i] = 1
	}
	return UF{set, count, n}
}

func (uf *UF) Find(a int) int {
	p := uf.set[a]
	if p != a {
		uf.set[a] = uf.Find(p)
	}
	return uf.set[a]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		// no change
		return false
	}

	if uf.count[pa] > uf.count[pb] {
		uf.set[pb] = pa
		uf.count[pa] += uf.count[pb]
	} else {
		uf.set[pa] = pb
		uf.count[pb] += uf.count[pa]
	}

	uf.size--
	return true
}

func (uf UF) Size() int {
	return uf.size
}
