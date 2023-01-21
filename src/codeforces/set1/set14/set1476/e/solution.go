package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n, m, k := readThreeNums(reader)

	P := make([]string, n)
	for i := 0; i < n; i++ {
		P[i], _ = reader.ReadString('\n')
		P[i] = P[i][:k]
	}

	S := make([]string, m)
	T := make([]int, m)

	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		S[i] = string(s[:k])
		readInt(s, k+1, &T[i])
	}

	res := solve(k, P, S, T)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	buf.WriteString("YES\n")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
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

func solve(k int, P []string, S []string, T []int) []int {
	K := 1 << k
	ii := make([]*Node, K)

	for i := 0; i < K; i++ {
		ii[i] = new(Node)
	}
	var buf bytes.Buffer
	for i := 0; i < len(P); i++ {
		cur := P[i]
		var mask int
		for j := 0; j < k; j++ {
			if cur[j] == '_' {
				mask |= 1 << j
			} else {
				buf.WriteByte(cur[j])
			}
		}
		AddString(ii[mask], buf.String(), i)
		buf.Reset()
	}

	n := len(P)
	que := make([]int, n)
	// at most K * n edges, when each mask contains all node
	var all int

	for i := 0; i < len(S); i++ {
		// lets enumerate the masks
		var p int
		for mask := 0; mask < K; mask++ {
			for j := 0; j < k; j++ {
				if (mask>>j)&1 == 0 {
					//should be same
					buf.WriteByte(S[i][j])
				}
			}
			res := Search(ii[mask], buf.String())
			buf.Reset()
			if len(res) > 0 {
				copy(que[p:], res)
				p += len(res)
			}
		}
		if p == 0 {
			return nil
		}
		all += p - 1
	}

	g := NewGraph(n, all)

	for i := 0; i < len(S); i++ {
		// lets enumerate the masks
		var p int
		for mask := 0; mask < K; mask++ {
			for j := 0; j < k; j++ {
				if (mask>>j)&1 == 0 {
					//should be same
					buf.WriteByte(S[i][j])
				}
			}
			res := Search(ii[mask], buf.String())
			buf.Reset()
			if len(res) > 0 {
				copy(que[p:], res)
				p += len(res)
			}
		}

		// if not found T[i] return nil
		pivot := -1
		for j := 0; j < p; j++ {
			if que[j] == T[i]-1 {
				pivot = j
				break
			}
		}
		if pivot < 0 {
			return nil
		}
		for j := 0; j < p; j++ {
			if j == pivot {
				continue
			}
			g.AddEdge(que[pivot], que[j])
		}
	}
	vis := make([]int, n)
	res := make([]int, 0, n)
	var topo func(u int) bool
	topo = func(u int) bool {
		if vis[u] == 1 {
			return false
		}
		vis[u]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 2 {
				continue
			}
			if !topo(v) {
				return false
			}
		}
		vis[u]++
		res = append(res, u+1)
		return true
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			if !topo(i) {
				return nil
			}
		}
	}

	reverse(res)

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	cur := 1
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type Node struct {
	children [26]*Node
	pos      []int
}

func AddString(root *Node, s string, p int) {
	node := root
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if node.children[x] == nil {
			node.children[x] = new(Node)
		}
		node = node.children[x]
	}
	if node.pos == nil {
		node.pos = make([]int, 0, 2)
	}
	node.pos = append(node.pos, p)
}

func Search(root *Node, s string) []int {
	node := root
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if node.children[x] == nil {
			return nil
		}
		node = node.children[x]
	}
	return node.pos
}
