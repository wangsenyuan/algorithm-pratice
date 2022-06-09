package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)

	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int) int {
	n := len(A)
	prev := make([]int, n+1)
	first := make(map[int]int)
	last := make(map[int]int)

	for i := 1; i <= n; i++ {
		cur := A[i-1]
		if _, ok := first[cur]; !ok {
			first[cur] = i
		}
		prev[i] = last[cur]
		last[cur] = i
	}

	id := make([]int, n+1)
	dp := make([]int, n+1)
	id[0] = 1

	tree := NewSegment(n + 1)

	M := make([]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		M[i] = make(map[int]int)
	}

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if prev[i] > 0 {
			tree.Update(prev[i], Node{first[A[i-1]], last[A[i-1]], 0})
		}
		tree.Update(i, Node{first[A[i-1]], last[A[i-1]], A[i-1]})

		if last[A[i-1]] == i {
			l := first[A[i-1]]
			r := last[A[i-1]]
			now := tree.Query(l, r+1)

			for now.last == i {
				if now.first == l {
					break
				}
				l = now.first
				now = tree.Query(l, r+1)
			}

			l = now.first

			if now.last == i {
				id[i] = id[l-1]
				now = tree.Query(id[i], i+1)

				if M[id[i]][now.xor] != 0 {
					dp[i] = max(dp[i], dp[M[id[i]][now.xor]]+1)
				} else if now.xor == 0 {
					dp[i] = max(dp[i], dp[id[i]-1]+1)
				}
				M[id[i]][now.xor] = i
			} else {
				id[i] = i + 1
			}
		} else {
			id[i] = i + 1
		}
	}

	return dp[n]
}

type Node struct {
	first int
	last  int
	xor   int
}

func (this *Node) Assign(that Node) {
	this.first = that.first
	this.last = that.last
	this.xor = that.xor
}

func (this *Node) Merge(a, b *Node) {
	this.first = min(a.first, b.first)
	this.last = max(a.last, b.last)
	this.xor = a.xor ^ b.xor
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Segment struct {
	nodes []*Node
	sz    int
}

const INF = 1 << 30

func NewSegment(n int) *Segment {
	nodes := make([]*Node, 2*n)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = new(Node)
		nodes[i].first = INF
		nodes[i].last = 0
		nodes[i].xor = 0
	}
	return &Segment{nodes, n}
}

func (tree *Segment) Update(p int, v Node) {
	p += tree.sz
	nodes := tree.nodes
	nodes[p].Assign(v)
	for p > 0 {
		nodes[p>>1].Merge(nodes[p], nodes[p^1])
		p >>= 1
	}
}

func (tree *Segment) Query(l int, r int) *Node {
	res := new(Node)
	res.first = INF
	res.last = 0
	res.xor = 0

	l += tree.sz
	r += tree.sz

	for l < r {
		if l&1 == 1 {
			res.Merge(res, tree.nodes[l])
			l++
		}
		if r&1 == 1 {
			r--
			res.Merge(res, tree.nodes[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
