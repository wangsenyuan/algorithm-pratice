package main

import (
	"bufio"
	"bytes"
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
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	head := readTree(reader)
	Q := readNum(reader)
	qs := make([]*Tree, Q)
	for i := 0; i < Q; i++ {
		qs[i] = readTree(reader)
	}
	res := solve(head, qs)
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		if res[i] {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
}

func readTree(reader *bufio.Reader) *Tree {
	n := readNum(reader)
	ages := make([]int, n)
	for i := 0; i < n; i++ {
		ages[i] = readNum(reader)
	}
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		s, _ := reader.ReadBytes('\n')
		edges[i] = make([]int, 3)
		pos := readInt(s, 0, &edges[i][0])
		pos = readInt(s, pos+1, &edges[i][1])
		if s[pos+2] == 'E' {
			edges[i][2] = 1
		}
	}
	return CreateTree(n, ages, edges)
}

type Tree struct {
	n     int
	ages  []int
	edges [][]int
}

func CreateTree(n int, ages []int, edges [][]int) *Tree {
	return &Tree{n, ages, edges}
}

func (t *Tree) String() string {
	n := t.n
	parent := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
		left[i] = -1
		right[i] = -1
	}

	for _, e := range t.edges {
		a, b, c := e[0]-1, e[1]-1, e[2]
		// M is 0, E is 1
		if c == 0 {
			left[a] = b
		} else {
			right[a] = b
		}
		parent[b] = a
	}

	return createString(n, left, right, parent, t.ages)
}

func solve(src *Tree, qs []*Tree) []bool {
	orig := src.String()

	n := len(orig)

	find := func(q *Tree) bool {
		pat := q.String()
		lps := kmp(pat)

		for i, j := 0, 0; i < n; {
			if orig[i] == pat[j] {
				i++
				j++
			}
			if j == len(pat) {
				return true
			}
			if i < n && orig[i] != pat[j] {
				if j > 0 {
					j = lps[j-1]
				} else {
					i++
				}
			}
		}

		return false
	}

	res := make([]bool, len(qs))

	for i, q := range qs {
		res[i] = find(q)
	}

	return res
}

func kmp(s string) []int {
	n := len(s)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		j := res[i-1]
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		res[i] = j
		if s[i] == s[j] {
			res[i]++
		}
	}
	return res
}

type ListNode struct {
	val  byte
	next *ListNode
	last *ListNode
}

func NewListNode(v int) *ListNode {
	node := new(ListNode)
	node.val = byte(v)
	node.last = node
	return node
}

type Node struct {
	left  *Node
	right *Node
	ln    *ListNode
	index int
}

func NewNode(index int) *Node {
	node := new(Node)
	node.index = index
	return node
}

func createString(n int, left []int, right []int, parent []int, age []int) string {
	var root int
	for i := 0; i < n; i++ {
		if parent[i] < 0 {
			root = i
			break
		}
	}
	var layers [][]*Node
	rootNode := NewNode(root)
	current := []*Node{rootNode}

	for len(current) > 0 {
		layers = append(layers, current)
		var next []*Node
		for i := 0; i < len(current); i++ {
			node := current[i]
			val := node.index
			lc := left[val]
			if lc != -1 {
				child := NewNode(lc)
				node.left = child
				next = append(next, child)
			}
			rc := right[val]
			if rc != -1 {
				child := NewNode(rc)
				node.right = child
				next = append(next, child)
			}
		}
		current = next
	}

	for i := len(layers) - 1; i >= 0; i-- {
		list := layers[i]
		for j := 0; j < len(list); j++ {
			current := list[j]
			cln := NewListNode(age[current.index])
			ln1 := NewListNode(int('z'))
			ln2 := NewListNode(int('z'))
			if current.left != nil {
				ln1 = current.left.ln
			}
			if current.right != nil {
				ln2 = current.right.ln
			}
			ln1.last.next = ln2
			ln2.last.next = cln
			current.ln = ln1
			current.ln.last = cln
		}
	}

	var ans bytes.Buffer

	link := rootNode.ln
	for link != nil {
		ans.WriteByte(link.val)
		link = link.next
	}
	return ans.String()
}
