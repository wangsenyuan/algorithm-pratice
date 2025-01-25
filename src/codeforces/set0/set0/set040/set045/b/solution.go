package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	p := readNNums(reader, n)
	v := readNNums(reader, m)
	b := readNNums(reader, m)
	return solve(n, m, p, v, b)
}

type pair struct {
	first  int
	second int
}

func solve(n int, m int, p []int, v []int, b []int) []int {
	fa := make([]int, n)
	adj := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		p[i]--
		if p[i] != i {
			fa[i] = p[i]
			adj[p[i]] = append(adj[p[i]], i)
			deg[p[i]]++
		}
	}

	que := make([]int, n)
	dist := make([]int, n)
	comp := make([]int, n)

	var head, tail int
	for u := 0; u < n; u++ {
		dist[u] = -1
		comp[u] = -1
		if deg[u] == 0 {
			que[head] = u
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		v := fa[u]
		deg[v]--
		if deg[v] == 0 {
			que[head] = v
			head++
		}
	}
	// head < n must hold
	// pos[u] = u在环上的位置
	head, tail = 0, 0
	pos := make([]int, n)
	var cycles []pair
	var id int

	tr := NewSegTree(n)
	at := make([]int, n)

	for u := 0; u < n; u++ {
		if deg[u] > 0 && comp[u] < 0 {
			v := u
			cur := id
			for comp[v] < 0 {
				tr.Update(id, id)
				pos[v] = id
				at[id] = v
				id++
				comp[v] = len(cycles)
				v = fa[v]
				que[head] = v
				head++
				dist[v] = 0
			}
			cycles = append(cycles, pair{cur, id})
		}
	}

	// 处理不在环上的部分
	for tail < head {
		u := que[tail]
		tail++
		for _, v := range adj[u] {
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	marked := make([]bool, n)

	var find func(u int, d int) (p int, cnt int)

	find = func(u int, d int) (p int, cnt int) {
		if comp[u] < 0 {
			if !marked[u] {
				cnt++
				marked[u] = true
			}
			if dist[u]-dist[fa[u]] > d {
				return fa[u], cnt
			}
			// dist[u] - dist[fa[u]] < d
			var sz int
			p, sz = find(fa[u], d-(dist[u]-dist[fa[u]]))
			cnt += sz
			fa[u] = p
			return
		}
		// u in cycle
		p = u
		// in_cycle[u] >= 0
		cycle := cycles[comp[u]]
		i := pos[u]
		if !marked[u] {
			marked[u] = true
			cnt++
			tr.Update(i, inf)
		}
		for d > 0 {
			j := tr.Get(i, cycle.second)
			if j == inf {
				d -= cycle.second - i
				if d < 0 {
					break
				}
				// 后面没有，检查前面
				j = tr.Get(cycle.first, cycle.second)
				i = cycle.first
			}

			d -= j - i
			if d < 0 {
				break
			}

			if !marked[at[j]] {
				cnt++
				marked[at[j]] = true
				tr.Update(j, inf)
			}
			i = j
		}
		return
	}

	var res int
	var ans []int

	for i := range m {
		a := (v[i] + res - 1 + n) % n
		_, res = find(a, b[i]-1)
		ans = append(ans, res)
	}

	return ans
}

type SegTree struct {
	arr []int
	sz  int
}

const inf = 1 << 30

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := range 2 * n {
		arr[i] = inf
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := inf
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
