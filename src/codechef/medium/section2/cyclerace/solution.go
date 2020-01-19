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

func readAtMostNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, 0, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n && x < len(scanner.Bytes()); i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}

		if x == len(scanner.Bytes()) {
			break
		}

		var num int
		x = readInt(scanner.Bytes(), x, &num)
		res = append(res, num)
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

	n, Q := readTwoNums(scanner)

	solver := NewSolver(n)

	for i := 0; i < Q; i++ {
		line := readAtMostNNums(scanner, 4)
		if line[0] == 1 {
			tm, j, sp := line[1], line[2]-1, line[3]
			solver.Change(tm, j, sp)
		} else {
			tm := line[1]
			res := solver.GetMax(tm)
			fmt.Println(res)
		}
	}
}

type Solver struct {
	n        int
	speed    []int64
	position []int64
	root     *Node
}

const MAXH = 1e9

func NewSolver(n int) Solver {
	speed := make([]int64, n)
	position := make([]int64, n)
	root := NewNode(1, MAXH)
	return Solver{n, speed, position, root}
}

func (solver *Solver) Change(t int, i int, sp int) {
	solver.position[i] += solver.speed[i] * int64(t)
	solver.position[i] -= int64(sp) * int64(t)
	solver.speed[i] = int64(sp)

	line := Line{-solver.speed[i], -solver.position[i]}

	solver.root.Update(line)
}

func (solver *Solver) GetMax(t int) int64 {
	T := int64(t)
	res := solver.root.GetMin(T)
	return -res
}

type Line struct {
	a, b int64
}

func (line Line) Eval(x int64) int64 {
	return line.a*x + line.b
}

type Node struct {
	l, r        int64
	line        Line
	left, right *Node
}

func NewNode(l, r int64) *Node {
	node := new(Node)
	node.l = l
	node.r = r
	return node
}

func (node *Node) pushDown() {
	if node.l < node.r && node.left == nil {
		mid := (node.l + node.r) / 2
		node.left = NewNode(node.l, mid)
		node.right = NewNode(mid+1, node.r)
	}
}

func (node *Node) Update(newLine Line) {
	node.pushDown()

	a, b := node.line, newLine

	if a.Eval(node.l) < b.Eval(node.l) {
		a, b = b, a
	}

	if a.Eval(node.r) >= b.Eval(node.r) {
		// b is totally down under a, no need to keep a any more
		node.line = b
		return
	}

	if node.l < node.r {
		mid := (node.l + node.r) / 2

		if a.Eval(mid) > b.Eval(mid) {
			// right has more less value after mid
			node.line = b
			node.right.Update(a)
		} else {
			node.line = a
			node.left.Update(b)
		}
	}
}

func (node *Node) GetMin(x int64) int64 {
	ans := node.line.Eval(x)

	mid := (node.l + node.r) / 2

	if mid < x && node.right != nil {
		return min(ans, node.right.GetMin(x))
	}

	if mid >= x && node.left != nil {
		return min(ans, node.left.GetMin(x))
	}

	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
