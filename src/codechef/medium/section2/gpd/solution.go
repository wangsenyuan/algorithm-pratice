package main

import (
	"bufio"
	"fmt"
	"os"
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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}

func readUint(bs []byte, from int, val *uint) int {
	i := from
	var res uint
	for i < len(bs) && bs[i] != ' ' {
		res = res*10 + uint(bs[i]-'0')
		i++
	}
	*val = res
	return i
}

func readNUints(scanner *bufio.Scanner, n int) []uint {
	res := make([]uint, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readUint(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readUintsUtilEnd(scanner *bufio.Scanner, hint int) []uint {
	res := make([]uint, 0, hint)
	x := 0
	scanner.Scan()
	for x < len(scanner.Bytes()) {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		if x < len(scanner.Bytes()) {
			var tmp uint
			x = readUint(scanner.Bytes(), x, &tmp)
			res = append(res, tmp)
		}
		x++
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	N, Q := readTwoNums(scanner)

	secondLine := readNUints(scanner, 2)
	R, RK := secondLine[0], secondLine[1]

	edges := make([][]uint, N-1)
	for i := 0; i < N-1; i++ {
		edges[i] = readNUints(scanner, 3)
	}
	queries := make([][]uint, Q)
	for i := 0; i < Q; i++ {
		queries[i] = readUintsUtilEnd(scanner, 4)
	}
	res := solve(N, R, RK, edges, Q, queries)

	for _, ans := range res {
		fmt.Printf("%d %d\n", ans[0], ans[1])
	}
}

func solve(N int, R uint, RK uint, edges [][]uint, Q int, queries [][]uint) [][]uint {
	forest := make(map[uint]*Trie)
	forest[R] = forest[R].Add(RK)

	outs := make(map[uint]map[uint]uint)

	for _, edge := range edges {
		u, v, k := edge[0], edge[1], edge[2]
		if _, found := outs[v]; !found {
			outs[v] = make(map[uint]uint)
		}
		outs[v][u] = k
	}

	var dfs func(u uint)

	dfs = func(u uint) {
		for v, k := range outs[u] {
			forest[v] = forest[u].Add(k)
			dfs(v)
		}
	}

	dfs(R)

	res := make([][]uint, 0, Q)

	var lastAnswer uint

	for i := 0; i < Q; i++ {
		query := queries[i]
		query[0] ^= lastAnswer
		if query[0] == 0 {
			v, u, k := query[1]^lastAnswer, query[2]^lastAnswer, query[3]^lastAnswer
			forest[u] = forest[v].Add(k)
		} else if query[0] == 1 {
			v, k := query[1]^lastAnswer, query[2]^lastAnswer
			a := forest[v].GetMin(k)
			b := forest[v].GetMax(k)
			res = append(res, []uint{a, b})
			lastAnswer = a ^ b
		} else {
			panic(fmt.Sprintf("some wrong before %v", query))
		}
	}

	return res
}

type Trie struct {
	children []*Trie // only 0 and 1
}

func newTrieNode(src *Trie) *Trie {
	dst := new(Trie)
	dst.children = make([]*Trie, 2)
	if src != nil {
		copy(dst.children, src.children)
	}
	return dst
}

func (root *Trie) Add(num uint) *Trie {

	var loop func(node *Trie, idx int) *Trie

	loop = func(node *Trie, idx int) *Trie {
		if idx < 0 {
			// at most 31
			return new(Trie)
		}

		node = newTrieNode(node)

		x := (num >> uint(idx)) & 1

		node.children[x] = loop(node.children[x], idx-1)

		return node
	}

	return loop(root, 31)
}

func (root *Trie) GetMax(key uint) uint {
	var res uint
	node := root
	for i := 31; i >= 0; i-- {
		x := (key >> uint(i)) & 1
		if node.children[1-x] != nil {
			res |= 1 << uint(i)
			node = node.children[1-x]
		} else {
			node = node.children[x]
		}
	}
	return res
}

func (root *Trie) GetMin(key uint) uint {
	var res uint
	node := root
	for i := 31; i >= 0; i-- {
		x := (key >> uint(i)) & 1
		if node.children[x] != nil {
			node = node.children[x]
		} else {
			res |= 1 << uint(i)
			node = node.children[1-x]
		}
	}

	return res
}
