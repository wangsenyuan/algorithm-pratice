package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	var n int
	for i := 0; i < len(s); i++ {
		if s[i] < '1' || s[i] > '9' {
			break
		}
		n++
	}
	x := readNum(reader)

	res := solve(s[:n], x)

	fmt.Println(res)
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

func solve(s string, x int) int {

	trie := NewTrie(N)

	var generate func(sum int)

	buf := make([]byte, 0, 100)

	generate = func(sum int) {
		if sum == x {
			var ok = true
			for l := 0; l < len(buf) && ok; l++ {
				var cur int
				for r := l; r < len(buf) && ok; r++ {
					cur += int(buf[r] - '0')
					if x%cur == 0 && cur != x {
						ok = false
					}
				}
			}
			if ok {
				trie.AddString(string(buf))
			}
			return
		}
		for j := 1; j <= min(x-sum, 9); j++ {
			n := len(buf)
			buf = append(buf, byte(j+'0'))
			generate(sum + j)
			buf = buf[:n]
		}
	}

	generate(0)

	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(trie.nodes)+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(trie.nodes); j++ {
			if dp[i][j] == INF {
				continue
			}
			dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			next := trie.Transition(j, int(s[i]-'1'))
			if !trie.nodes[next].leaf {
				dp[i+1][next] = min(dp[i+1][next], dp[i][j])
			}
		}
	}
	var res = INF

	for j := 0; j < len(trie.nodes); j++ {
		res = min(res, dp[len(s)][j])
	}

	return res
}

const AL = 9
const N = 5000
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
		c := int(s[i] - '1')
		if nodes[v].next[c] == -1 {
			nodes[v].next[c] = len(nodes)
			cur := NewNode()
			nodes = append(nodes, cur)
			cur.parent = v
			cur.parentByte = c
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
