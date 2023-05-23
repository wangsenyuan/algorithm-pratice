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

	tc := readNum(reader)

	readPeople := func(n int, att [][]int, like [][]int) {
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var a, b, m int
			pos := readInt(s, 0, &a)
			pos = readInt(s, pos+1, &b)
			pos = readInt(s, pos+1, &m)
			att[i] = []int{a, b}
			like[i] = make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(s, pos+1, &like[i][j])
			}
		}
	}

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		B, G, L := readThreeNums(reader)
		boys := make([][]int, B)
		boys_like := make([][]int, B)
		readPeople(B, boys, boys_like)
		girls := make([][]int, G)
		girls_like := make([][]int, G)
		readPeople(G, girls, girls_like)
		res := solve(L, B, G, boys, girls, boys_like, girls_like)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(L int, B int, G int, boys_att [][]int, girls_att [][]int, boys_like [][]int, girls_like [][]int) []int {
	hit := make([][]int, B)
	for i := 0; i < B; i++ {
		hit[i] = make([]int, G)
	}

	for i := 0; i < B; i++ {
		for _, j := range boys_like[i] {
			if j < G {
				hit[i][j]++
			}
		}
	}
	for j := 0; j < G; j++ {
		for _, i := range girls_like[j] {
			if i < B {
				hit[i][j]++
			}
		}
	}
	g := NewGraph(B+G, B*G)

	for i := 0; i < B; i++ {
		for j := 0; j < G; j++ {
			if hit[i][j] == 2 {
				g.AddEdge(i, B+j)
			}
		}
	}
	type Event struct {
		id   int
		time int
	}

	events := make([]Event, 0, 2*len(boys_att)+2*len(girls_att))

	for i, cur := range boys_att {
		a, b := cur[0], cur[1]
		events = append(events, Event{i, a})
		events = append(events, Event{i, b})
	}

	for i, cur := range girls_att {
		a, b := cur[0], cur[1]
		events = append(events, Event{B + i, a})
		events = append(events, Event{B + i, b})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time < events[j].time
	})

	active := make([]bool, B+G)

	pair := make([][]int, 2)
	pair[0] = make([]int, B)
	pair[1] = make([]int, G)
	for i := 0; i < 2; i++ {
		for j := 0; j < len(pair[i]); j++ {
			pair[i][j] = -1
		}
	}

	seen := make([]bool, G)

	var dfs func(u int) bool

	dfs = func(boy int) bool {
		if !active[boy] {
			return false
		}
		for i := g.node[boy]; i > 0; i = g.next[i] {
			girl := g.to[i] - B
			if seen[girl] || pair[1][girl] == boy {
				continue
			}
			seen[girl] = true
			if pair[1][girl] == -1 || dfs(pair[1][girl]) {
				pair[1][girl] = boy
				pair[0][boy] = girl
				return true
			}
		}
		return false
	}

	aug := func() bool {
		for i := 0; i < G; i++ {
			seen[i] = !active[i+B]
		}
		for i := 0; i < B; i++ {
			if pair[0][i] < 0 && dfs(i) {
				return true
			}
		}
		return false
	}
	ans := make([]int, min(B, G)+1)
	var prev int
	var cnt int
	for _, cur := range events {
		ans[cnt] += cur.time - prev

		if active[cur.id] {
			// a leave event
			i := 0
			j := cur.id
			if j >= B {
				i = 1
				j -= B
			}

			if pair[i][j] != -1 {
				pair[1^i][pair[i][j]] = -1
				pair[i][j] = -1
				cnt--
			}
		}

		active[cur.id] = !active[cur.id]

		if aug() {
			cnt++
		}

		prev = cur.time
	}

	ans[cnt] += L - prev

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.to[g.cur] = v
	g.node[u] = g.cur
}
