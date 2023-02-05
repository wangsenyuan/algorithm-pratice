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

	n, m := readTwoNums(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 3)
	}

	res := solve(n, Q)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const H = 31

func solve(n int, E [][]int) []int {
	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = (1 << H) - 1
	}

	solve_bit := func(pt int) bool {
		for u := 0; u < n; u++ {
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if get_bit(w, pt) == 0 {
					// 如果 A[u] | A[v] = 0
					set_bit(&ans[u], pt, 0)
				} else if get_bit(ans[u], pt) == 0 && get_bit(ans[v], pt) == 0 {
					// 如果结果为1, 但是 A[u] | A[v] = 0
					return false
				}
			}
		}

		for u := 0; u < n; u++ {
			if get_bit(ans[u], pt) == 1 {
				// try set it as 0
				set_bit(&ans[u], pt, 0)
				for i := g.node[u]; i > 0; i = g.next[i] {
					v := g.to[i]
					w := g.val[i]
					if get_bit(ans[v], pt) == 0 && get_bit(w, pt) == 1 {
						set_bit(&ans[u], pt, 1)
						break
					}
				}
			}
		}
		return true
	}

	for i := 0; i < H; i++ {
		if !solve_bit(i) {
			return nil
		}
	}

	return ans
}

func get_bit(num int, p int) int {
	return (num >> p) & 1
}

func set_bit(num *int, p int, v int) {
	if v == 0 {
		*num &= ^(1 << p)
	} else {
		*num |= (1 << p)
	}
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
