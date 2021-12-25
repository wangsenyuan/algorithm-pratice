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

	K, Q := readTwoNums(scanner)

	solver := NewSolver(K)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		line := scanner.Bytes()
		if line[0] == '!' {
			var l, r, x int
			pos := readInt(line, 2, &l) + 1
			pos = readInt(line, pos, &r) + 1
			readInt(line, pos, &x)
			solver.AddRangeValue(l, r, x)
		} else {
			var l, r int
			pos := readInt(line, 2, &l) + 1
			pos = readInt(line, pos, &r) + 1
			ans := solver.GetCount(l, r)
			fmt.Println(ans)
		}
	}
}

type Solver struct {
	root    *Node
	factors []int
}

const MAX_X = 100007

func NewSolver(K int) Solver {
	root := NewNode(1, MAX_X)

	factors := getPrimeFactors(K)

	return Solver{root, factors}
}

func (solver *Solver) AddRangeValue(l, r, x int) {
	var v int

	for i := 0; i < len(solver.factors); i++ {
		f := solver.factors[i]
		if x%f == 0 {
			v |= 1 << uint(i)
		}
	}
	if v > 0 {
		solver.root.SetValue(l, r, v)
	}
}

func (solver Solver) GetCount(l, r int) int {
	v := solver.root.GetValue(l, r)
	var count int
	for v > 0 {
		count += v & 1
		v >>= 1
	}
	return count
}

func getPrimeFactors(K int) []int {
	fs := make([]int, 0, 32)

	f := 2
	for {
		if f*f > K {
			if K > 1 {
				fs = append(fs, K)
				break
			}
		}

		if K%f == 0 {
			fs = append(fs, f)
			for K%f == 0 {
				K /= f
			}
		}

		f++
	}

	return fs
}

type Node struct {
	start int
	end   int
	v     int
	left  *Node
	right *Node
	set   bool
}

func (node *Node) getMid() int {
	return (node.start + node.end) / 2
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	return node
}

func (node *Node) pushDown() {
	if node.start < node.end && node.left == nil {
		mid := node.getMid()
		node.left = NewNode(node.start, mid)
		node.right = NewNode(mid+1, node.end)
	}
}

func (node *Node) SetValue(a, b, v int) {
	if node == nil {
		return
	}
	node.pushDown()

	if a > node.end || b < node.start {
		// out of range
		return
	}

	if node.set {
		// this range already set by other values
		return
	}

	// have overlap
	if a <= node.start && node.end <= b {
		// totally in the range
		node.set = true
		node.v = v
		return
	}

	// mid := node.getMid()

	node.left.SetValue(a, b, v)
	node.right.SetValue(a, b, v)

	if node.left.set && node.right.set {
		node.set = true
		node.v = node.left.v | node.right.v
	}
}

func (node *Node) GetValue(a, b int) int {
	if node == nil {
		return 0
	}
	if a > node.end || b < node.start {
		// out of range
		return 0
	}
	if a <= node.start && node.end <= b {
		if node.set {
			return node.v
		}
	}

	if node.left == nil {
		return 0
	}
	return node.left.GetValue(a, b) | node.right.GetValue(a, b)
}
