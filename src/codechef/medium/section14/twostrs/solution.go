package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var ff bytes.Buffer
	for tc > 0 {
		tc--
		A, _ := reader.ReadString('\n')
		A = format(A)
		B, _ := reader.ReadString('\n')
		B = format(B)

		n := readNum(reader)
		S := make([]string, n)
		bb := make([]int, n)
		for i := 0; i < n; i++ {
			buf, _ := reader.ReadBytes('\n')
			var j int
			for j < len(buf) && buf[j] != ' ' {
				j++
			}
			// buf[j] = ' '
			S[i] = string(buf[:j])
			readInt(buf, j+1, &bb[i])
		}
		res := solve(A, B, n, S, bb)

		ff.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(ff.String())
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

func format(s string) string {
	buf := []byte(s)

	for i := 0; i < len(buf); i++ {
		if buf[i] < 'a' || buf[i] > 'z' {
			return string(buf[:i])
		}
	}

	return s
}

func solve(A, B string, N int, S []string, BB []int) int64 {
	trie := NewTrie(NODE_CNT)

	for i := 0; i < N; i++ {
		trie.AddString(S[i], BB[i])
	}

	dp := make([]int64, len(A))
	var cur int
	for i := 0; i < len(A); i++ {
		cur = trie.Transition(cur, int(A[i]-'a'))
		dp[i] = trie.nodes[cur].score
		if i > 0 {
			dp[i] += dp[i-1]
		}
	}

	cur = 0

	fp := make([]int64, len(B))

	for i := 0; i < len(B); i++ {
		cur = trie.Transition(cur, int(B[i]-'a'))
		fp[i] = trie.nodes[cur].score
	}

	for i := len(B) - 2; i >= 0; i-- {
		fp[i] += fp[i+1]
	}

	var best int64

	cur = 0

	for i := 0; i < len(A); i++ {
		cur = trie.Transition(cur, int(A[i]-'a'))
		for j := 0; j < len(B); j++ {
			tmp := dp[i]
			node := cur
			for k := j; k <= j+24 && k < len(B); k++ {
				node = trie.Transition(node, int(B[k]-'a'))
				tmp += trie.nodes[node].score
			}
			if j+25 < len(B) {
				tmp += fp[j+25]
			}
			if best < tmp {
				best = tmp
			}
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const AL = 26
const NODE_CNT = 2010
const INF = 1000000000

type Node struct {
	next       [AL]int
	parent     int
	parentByte int
	suffixLink int
	transition [AL]int
	leaf       bool
	score      int64
}

func NewNode() *Node {
	node := new(Node)
	for i := 0; i < AL; i++ {
		node.next[i] = -1
		node.transition[i] = -1
	}
	node.suffixLink = -1
	node.parent = -1
	node.score = 0
	return node
}

type Trie struct {
	nodes []*Node
}

func NewTrie(sizeHint int) *Trie {
	trie := new(Trie)
	trie.nodes = make([]*Node, 0, sizeHint+1)
	trie.nodes = append(trie.nodes, NewNode())
	return trie
}

func (trie *Trie) AddString(s string, score int) {
	var v int
	nodes := trie.nodes
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		if c < 0 || c >= AL {
			break
		}
		if nodes[v].next[c] == -1 {
			nodes[v].next[c] = len(nodes)
			cur := NewNode()
			cur.parent = v
			cur.parentByte = c
			nodes = append(nodes, cur)
		}
		v = nodes[v].next[c]
	}
	nodes[v].score += int64(score)
	nodes[v].leaf = true
	trie.nodes = nodes
}

func (trie *Trie) Transition(v int, c int) int {
	if trie.nodes[v].transition[c] == -1 {
		if trie.nodes[v].next[c] != -1 {
			trie.nodes[v].transition[c] = trie.nodes[v].next[c]
		} else {
			if v == 0 {
				trie.nodes[v].transition[c] = 0
			} else {
				trie.nodes[v].transition[c] = trie.Transition(trie.GetLink(v), c)
			}
		}
	}
	return trie.nodes[v].transition[c]
}

func (trie *Trie) GetLink(v int) int {
	if trie.nodes[v].suffixLink == -1 {
		if v == 0 || trie.nodes[v].parent == 0 {
			trie.nodes[v].suffixLink = 0
		} else {
			trie.nodes[v].suffixLink =
				trie.Transition(trie.GetLink(trie.nodes[v].parent), trie.nodes[v].parentByte)
		}
	}

	return trie.nodes[v].suffixLink
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
