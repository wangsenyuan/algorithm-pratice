package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
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

	Q := readNum(scanner)

	A := make([]byte, Q)
	B := make([]int, Q)
	C := make([]interface{}, Q)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		A[i] = scanner.Bytes()[0]

		pos := readInt(scanner.Bytes(), 2, &B[i])
		if A[i] == '?' {
			var c int
			readInt(scanner.Bytes(), pos+1, &c)
			C[i] = c
		} else {
			C[i] = string(scanner.Bytes()[pos+1:])
		}
	}
	res := solve(Q, A, B, C)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(Q int, A []byte, B []int, C []interface{}) []string {
	var root *Treap

	res := make([]string, 0, Q)

	var buf bytes.Buffer
	for i := 0; i < Q; i++ {
		if A[i] == '+' {
			pos := B[i]
			s := C[i].(string)
			for j := 0; j < len(s); j++ {
				root = root.Insert(pos+j+1, s[j])
			}
		} else {
			pos := B[i]
			ln := C[i].(int)
			buf.Reset()
			for j := 0; j < ln; j++ {
				buf.WriteByte(root.Get(pos + j))
			}
			res = append(res, buf.String())
		}
	}

	return res
}

type Treap struct {
	val         byte
	priority    int
	nrOfNodes   int
	left, right *Treap
}

func countNodes(node *Treap) int {
	if node == nil {
		return 0
	}
	return node.nrOfNodes
}

func update(node *Treap) {
	node.nrOfNodes = 1 + countNodes(node.left) + countNodes(node.right)
}

func rotateRight(node *Treap) *Treap {
	aux := node.left
	node.left = aux.right
	aux.right = node

	update(node)
	update(aux)
	return aux
}

func rotateLeft(node *Treap) *Treap {
	aux := node.right
	node.right = aux.left
	aux.left = node
	update(node)
	update(aux)
	return aux
}

func balance(node *Treap) *Treap {
	if node.left != nil && node.left.priority > node.priority {
		node = rotateRight(node)
	}
	if node.right != nil && node.right.priority > node.priority {
		node = rotateLeft(node)
	}
	update(node)
	return node
}

func (root *Treap) Insert(position int, val byte) *Treap {
	priority := rand.Int()
	var loop func(node *Treap, pos int) *Treap

	loop = func(node *Treap, pos int) *Treap {
		if node == nil {
			node = new(Treap)
			node.priority = priority
			node.val = val
			node.nrOfNodes = 1
			return node
		}
		if pos <= countNodes(node.left)+1 {
			node.left = loop(node.left, pos)
		} else {
			node.right = loop(node.right, pos-countNodes(node.left)-1)
		}
		return balance(node)
	}
	return loop(root, position)
}

func (root *Treap) Get(kth int) byte {
	var loop func(node *Treap, kth int) byte

	loop = func(node *Treap, kth int) byte {
		if node == nil {
			return '+'
		}
		if countNodes(node.left)+1 == kth {
			return node.val
		}
		if kth < countNodes(node.left)+1 {
			return loop(node.left, kth)
		}
		return loop(node.right, kth-countNodes(node.left)-1)
	}

	return loop(root, kth)
}

func solve1(Q int, A []byte, B []int, C []interface{}) []string {
	var N int
	qs := make([]*Query, 0, Q)
	for i := 0; i < Q; i++ {
		pos := B[i]
		if A[i] == '+' {
			s := C[i].(string)
			for j := len(s) - 1; j >= 0; j-- {
				q := Query{label: i, kind: '+', pos: pos + 1, c: s[j]}
				qs = append(qs, &q)
			}
			N += len(s)
		} else {
			ln := C[i].(int)
			for j := 0; j < ln; j++ {
				q := Query{label: i, kind: '?', pos: pos + j}
				qs = append(qs, &q)
			}
		}
	}
	st := initTree(N + 1)
	c := make([]byte, N+1)
	for i := len(qs) - 1; i >= 0; i-- {
		q := qs[i]
		if q.kind == '?' {
			q.pos = st.Query(q.pos)
		} else {
			q.pos = st.Query(q.pos)
			c[q.pos] = q.c
			st.Update(q.pos)
		}
	}

	res := make([]string, 0, len(qs))
	var buf bytes.Buffer
	for i := 0; i < len(qs); i++ {
		q := qs[i]
		if q.kind == '?' {
			buf.WriteByte(c[q.pos])
			if i == len(qs)-1 || q.label != qs[i+1].label {
				res = append(res, buf.String())
				buf.Reset()
			}
		}
	}

	return res
}

type Query struct {
	label int
	pos   int
	kind  byte
	c     byte
}

type SegTree struct {
	tree []int
	size int
}

func initTree(n int) *SegTree {
	tree := make([]int, 4*n)
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		tree[i] = end - start + 1
		if start == end {
			return
		}
		mid := (start + end) / 2
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
	}

	loop(0, 1, n)
	return &SegTree{tree: tree, size: n}
}

func (st SegTree) Query(kth int) int {
	tree := st.tree
	var loop func(i int, start, end int) int
	loop = func(i int, start, end int) int {
		if start == end {
			return start
		}
		mid := (start + end) / 2
		if kth > tree[2*i+1] {
			kth -= tree[2*i+1]
			return loop(2*i+2, mid+1, end)
		}

		return loop(2*i+1, start, mid)
	}

	return loop(0, 1, st.size)
}

func (st *SegTree) Update(pos int) {
	tree := st.tree
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if pos < start || pos > end {
			return
		}
		tree[i]--
		if start == end {
			return
		}
		mid := (start + end) / 2
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
	}
	loop(0, 1, st.size)
}
