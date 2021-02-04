package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	A := make([]string, n)
	for i := 0; i < n; i++ {
		A[i], _ = reader.ReadString('\n')
		A[i] = normalize(A[i])
	}
	L, R, X := make([]int, q), make([]int, q), make([]string, q)
	for i := 0; i < q; i++ {
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &L[i]) + 1
		pos = readInt(s, pos, &R[i]) + 1
		X[i] = normalize(string(s[pos:]))
	}
	res := solve(n, q, A, L, R, X)
	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {
	buf := []byte(s)
	var n int
	for n < len(buf) {
		if buf[n] != '0' && buf[n] != '1' {
			break
		}
		n++
	}
	return string(buf[:n])
}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

func solve(N, Q int, A []string, L []int, R []int, X []string) []int {
	var maxLen int

	for _, a := range A {
		maxLen = max(maxLen, len(a))
	}

	nodes := make([]*Node, maxLen+1)

	// just create a max len of '0', and insert
	nodes[0] = NewNode()
	for i := 1; i <= maxLen; i++ {
		nodes[i] = NewNode()
		nodes[i-1].children[0] = nodes[i]
	}

	for i := 0; i < N; i++ {
		pad := maxLen - len(A[i])
		Insert(nodes[pad], A[i], pad, i+1)
	}

	dfs(nodes[0])

	ans := make([]int, Q)

	for i := 0; i < Q; i++ {
		l, r, s := L[i], R[i], X[i]
		var pad int
		if len(s) > maxLen {
			s = s[len(s)-maxLen:]
		} else {
			pad = maxLen - len(s)
		}
		ans[i] = Search(nodes[0], l, r, s, pad)
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	children [2]*Node
	next     *Node
	depth    int
	v        []int
}

func NewNode() *Node {
	node := new(Node)
	node.v = make([]int, 0, 2)
	node.depth = 0
	return node
}

func (node *Node) AddIndex(i int) {
	node.v = append(node.v, i)
}

func Insert(root *Node, s string, dep int, idx int) {
	cur := root
	for i := 0; i < len(s); i++ {
		pos := int(s[i] - '0')
		if cur.children[pos] == nil {
			cur.children[pos] = NewNode()
			cur.children[pos].depth = dep + i + 1
		}
		cur = cur.children[pos]
	}
	cur.AddIndex(idx)
}

func dfs(node *Node) {
	if node.children[0] == nil && node.children[1] == nil {
		node.next = node
		return
	}
	if node.children[0] != nil {
		node.children[0].depth = node.depth + 1
		dfs(node.children[0])
	}
	if node.children[1] != nil {
		node.children[1].depth = node.depth + 1
		dfs(node.children[1])
	}

	if node.children[0] == nil || node.children[0].next == nil || len(node.children[0].next.v) == 0 {
		node.children[0] = nil
		if node.children[1] == nil {
			node.next = node
		} else {
			node.next = node.children[1].next
		}
		return
	}

	if node.children[1] == nil || node.children[1].next == nil || len(node.children[1].next.v) == 0 {
		node.children[1] = nil
		if node.children[0] == nil {
			node.next = node
		} else {
			node.next = node.children[0].next
		}
		return
	}

	node.v = merge(node.children[0].next.v, node.children[1].next.v)
	node.next = node
}

func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	var i, j, k int
	for i < len(a) || j < len(b) {
		if i == len(a) || j < len(b) && a[i] > b[j] {
			c[k] = b[j]
			j++
		} else {
			c[k] = a[i]
			i++
		}
		k++
	}
	return c
}

func Search(node *Node, L, R int, S string, pad int) int {
	cur := node
	for cur != nil {
		if cur.children[0] == nil && cur.children[1] == nil {
			break
		}
		var index int
		if cur.depth < pad {
			index = 1
		} else {
			index = 1 - int(S[cur.depth-pad]-'0')
		}
		if cur.children[index] == nil {
			cur = cur.children[1-index].next
		} else {
			V := cur.children[index].next.v
			pos := sort.Search(len(V), func(p int) bool {
				return V[p] >= L
			})
			if pos == len(V) || V[pos] > R {
				cur = cur.children[1-index].next
			} else {
				cur = cur.children[index].next
			}
		}
	}

	pos := sort.Search(len(cur.v), func(p int) bool {
		return cur.v[p] >= L
	})
	return cur.v[pos]
}
