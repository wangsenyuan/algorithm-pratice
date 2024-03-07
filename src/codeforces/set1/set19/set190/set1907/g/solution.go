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

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		lights := readString(reader)
		a := readNNums(reader, n)
		ok, res := solve(lights, a)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d", res[i]))
				if i+1 < len(res) {
					buf.WriteByte(' ')
				} else {
					buf.WriteByte('\n')
				}
			}
		}
	}

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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(lights string, a []int) (bool, []int) {
	n := len(lights)
	deg := make([]int, n)

	for i := 0; i < n; i++ {
		a[i]--
		deg[a[i]]++
	}

	que := make([]int, n)
	var front, tail int

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[front] = i
			front++
		}
	}

	on := make([]bool, n)
	for i := 0; i < n; i++ {
		on[i] = lights[i] == '1'
	}

	var res []int
	for tail < front {
		u := que[tail]
		tail++
		v := a[u]
		if on[u] {
			res = append(res, u)
			on[u] = !on[u]
			on[v] = !on[v]
		}
		deg[v]--
		if deg[v] == 0 {
			que[front] = v
			front++
		}
	}

	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		if on[i] && !vis[i] {
			x := i
			var p []int
			var ps []bool
			var cnt int
			for !vis[x] {
				p = append(p, x)
				ps = append(ps, on[x])
				if on[x] {
					cnt++
				}
				vis[x] = true
				x = a[x]
			}
			if cnt%2 == 1 {
				return false, nil
			}
			k := len(p)
			p = append(p, x)
			ps = append(ps, on[x])
			sp := make([]bool, len(ps))
			copy(sp, ps)
			var v1 []int
			for j := 0; j < k; j++ {
				if j == 0 || sp[j] {
					v1 = append(v1, p[j])
					sp[j] = !sp[j]
					sp[j+1] = !sp[j+1]
				}
			}
			var v2 []int
			for j := 0; j < k; j++ {
				if j != 0 && ps[j] {
					v2 = append(v2, p[j])
					ps[j] = !ps[j]
					ps[j+1] = !ps[j+1]
				}
			}
			if len(v1) > len(v2) {
				v1 = v2
			}
			res = append(res, v1...)
		}
	}
	for i := 0; i < len(res); i++ {
		res[i]++
	}
	return true, res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
