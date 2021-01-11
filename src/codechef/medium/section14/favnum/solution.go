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

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadBytes('\n')
	var pos int
	var L, R, K uint64
	pos = readUint64(s, 0, &L) + 1
	pos = readUint64(s, pos, &R) + 1
	pos = readUint64(s, pos, &K) + 1
	var n int
	readInt(s, pos, &n)
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i], _ = reader.ReadString('\n')
	}

	res := solve(int64(L), int64(R), int64(K), n, S)

	fmt.Println(res)
}

func solve(L, R, K int64, N int, S []string) string {
	trie := NewTrie(NODE_CNT)

	for _, s := range S {
		trie.AddString(s)
	}

	dp := make([][][]int64, len(trie.nodes))

	num := make([]int, 100)

	var doit func(n int, d int, f int) int64

	doit = func(n int, d int, f int) int64 {
		if d == -1 {
			// f & 1 == 1 means found
			return int64(f & 1)
		}
		if dp[n][d][f] >= 0 {
			return dp[n][d][f]
		}
		dp[n][d][f] = 0
		mx := 9
		if f&2 == 0 {
			//still tight
			mx = num[d]
		}

		for i := 0; i <= mx; i++ {
			var nf int
			if i != num[d] {
				nf |= 2
			}
			next := trie.Transition(n, i)
			if f&1 > 0 || trie.nodes[next].leaf {
				dp[n][d][f] += doit(0, d-1, f|nf|1)
				continue
			}
			dp[n][d][f] += doit(next, d-1, f|nf)
		}

		return dp[n][d][f]
	}

	reset := func() {
		for i := 0; i < len(dp); i++ {
			dp[i] = make([][]int64, 20)
			for j := 0; j < 20; j++ {
				dp[i][j] = make([]int64, 4)
				for f := 0; f < 4; f++ {
					dp[i][j][f] = -1
				}
			}
		}
	}

	count := func(x int64) int64 {
		reset()
		var i int
		for x > 0 {
			num[i] = int(x % 10)
			i++
			x /= 10
		}
		return doit(0, i-1, 0)
	}

	K += count(L - 1)
	D := count(R)

	if D < K {
		return "no such number"
	}

	for L < R {
		mid := (L + R) / 2
		D = count(mid)
		if D >= K {
			R = mid
		} else {
			L = mid + 1
		}
	}
	var i int
	for R > 0 {
		num[i] = int(R % 10)
		i++
		R /= 10
	}
	var buf bytes.Buffer

	for i > 0 {
		buf.WriteByte(byte('0' + num[i-1]))
		i--
	}
	return buf.String()
}

const AL = 10
const NODE_CNT = 2010
const INF = 1000000000

type Node struct {
	next       [AL]int
	parent     int
	parentByte int
	suffixLink int
	transition [AL]int
	leaf       bool
}

func NewNode() *Node {
	node := new(Node)
	for i := 0; i < AL; i++ {
		node.next[i] = -1
		node.transition[i] = -1
	}
	node.suffixLink = -1
	node.parent = -1
	node.leaf = false
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

func (trie *Trie) AddString(s string) {
	var v int
	nodes := trie.nodes
	for i := 0; i < len(s); i++ {
		c := int(s[i] - '0')
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
