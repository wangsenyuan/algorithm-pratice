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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && bs[x] == ' ' {
			x++
		}
		x = readInt(bs, x, &res[i])
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)
	arr := readNNums(scanner, n)

	solver := NewSolver(arr)

	input := make([]int, 4)

	for m > 0 {
		m--
		scanner.Scan()
		bs := scanner.Bytes()

		if bs[0] == '1' {
			fillNNums(bs, 3, input)

			res := solver.Query(input[1], input[2])
			fmt.Printf("%d %d\n", res.num, res.cnt)
		} else {
			fillNNums(bs, 4, input)
			solver.Modify(input[1], input[2], input[3])
		}
	}
}

type Solver struct {
	arr    []int
	blocks []*Block
}

const BLOCK_SIZE = 517

func NewSolver(arr []int) *Solver {
	n := len(arr)

	blocks := make([]*Block, (n+BLOCK_SIZE-1)/BLOCK_SIZE)

	for i := 0; i < n; {
		j := min(n, i+BLOCK_SIZE)
		block := NewBlock(arr, i, j-1)
		blocks[i/BLOCK_SIZE] = block
		i = j
	}

	return &Solver{arr, blocks}
}

func (this *Solver) Modify(l, r, k int) {
	l--
	r--
	li := l / BLOCK_SIZE
	ri := r / BLOCK_SIZE
	if li == ri {
		// same block
		for i := l; i <= r; i++ {
			this.arr[i] ^= k
		}
		this.blocks[li].Rebuild()
		return
	}
	for i := l; i <= this.blocks[li].right; i++ {
		this.arr[i] ^= k
	}
	this.blocks[li].ok = false
	for i := this.blocks[ri].left; i <= r; i++ {
		this.arr[i] ^= k
	}
	this.blocks[ri].ok = false

	for i := li + 1; i < ri; i++ {
		this.blocks[i].val ^= k
	}
}

func (this *Solver) Query(l, r int) Result {
	l--
	r--
	li := l / BLOCK_SIZE
	ri := r / BLOCK_SIZE
	result := Result{1 << 16, 0}

	if li == ri {
		for i := l; i <= r; i++ {
			result = result.Merge(Result{this.arr[i] ^ this.blocks[i/BLOCK_SIZE].val, 1})
		}
		return result
	}

	for i := l; i <= this.blocks[li].right; i++ {
		result = result.Merge(Result{this.arr[i] ^ this.blocks[i/BLOCK_SIZE].val, 1})
	}

	for i := this.blocks[ri].left; i <= r; i++ {
		result = result.Merge(Result{this.arr[i] ^ this.blocks[i/BLOCK_SIZE].val, 1})
	}

	for i := li + 1; i < ri; i++ {
		if !this.blocks[i].ok {
			this.blocks[i].Rebuild()
		}
		result = result.Merge(this.blocks[i].FindResult())
	}
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Node struct {
	children []*Node
	cnt      int
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 2)
	node.cnt = 0
	return node
}

func (node *Node) Add(x int, i int) {
	if i < 0 {
		node.cnt++
		return
	}
	d := (x >> uint(i)) & 1
	if node.children[d] == nil {
		node.children[d] = NewNode()
	}
	node.children[d].Add(x, i-1)
}

type Result struct {
	num int
	cnt int
}

func (this Result) Merge(that Result) Result {
	if this.num < that.num {
		return this
	} else if this.num > that.num {
		return that
	}
	return Result{this.num, this.cnt + that.cnt}
}

func findResult(node *Node, K int) Result {
	cur := node
	var num int
	for i := 15; i >= 0; i-- {
		d := (K >> uint(i)) & 1
		if cur.children[d] == nil {
			num |= (1 << uint(i))
			cur = cur.children[1-d]
			continue
		}
		cur = cur.children[d]
	}

	return Result{num, cur.cnt}
}

type Block struct {
	arr         []int
	left, right int
	ok          bool
	val         int
	tree        *Node
}

func NewBlock(arr []int, left, right int) *Block {
	block := new(Block)
	block.arr = arr
	block.left = left
	block.right = right
	block.tree = NewNode()
	block.val = 0
	block.ok = false
	return block
}

func (this *Block) Rebuild() {
	for i := this.left; i <= this.right; i++ {
		this.tree.Add(this.arr[i], 15)
	}
	this.ok = true
}

func (this Block) FindResult() Result {
	return findResult(this.tree, this.val)
}
