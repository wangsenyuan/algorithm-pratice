package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, q := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(n, E, Q)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const H = 30

func solve(n int, E [][]int, Q [][]int) []int {

	comp := make([][]int, H)
	for i := 0; i < H; i++ {
		comp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			comp[i][j] = -1
		}
	}

	dsu := NewDSU(n)

	or_mst := func(cur int) int {
		var ans int
		var bad int
		for i := H - 1; i >= 0; i-- {
			bad |= 1 << i
			comps := n
			dsu.Reset()

			for _, e := range E {
				u, v, w := e[0]-1, e[1]-1, e[2]
				if cur >= 0 {
					if w != 0 && comp[cur][u] != comp[cur][v] {
						continue
					}
				}
				if w&bad != 0 {
					continue
				}
				if dsu.Union(u, v) {
					comps--
				}
			}
			if comps == 1 {
				continue
			}
			// i不设置时，会产生多个分组
			bad ^= 1 << i
			ans ^= 1 << i

			if cur < 0 && comps == 2 {
				for u := 0; u < n; u++ {
					p := dsu.Find(u)
					comp[i][u] = p
				}
			}
		}

		return ans
	}

	mst := make([]int, H)

	init_ans := or_mst(-1)
	m := len(E)
	for i := 0; i < H; i++ {
		if comp[i][0] == -1 {
			continue
		}
		u := 0
		v := 1
		for v < n && comp[i][u] == comp[i][v] {
			v++
		}
		E = append(E, []int{u + 1, v + 1, 0})
		mst[i] = or_mst(i)
		E = E[:m]
	}

	res := make([]int, len(Q))

	for i, q := range Q {
		u, v := q[0]-1, q[1]-1
		res[i] = init_ans
		for j := H - 1; j >= 0; j-- {
			if comp[j][u] != comp[j][v] {
				res[i] = mst[j]
				break
			}
		}
	}

	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	d := new(DSU)
	d.arr = make([]int, n)
	d.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		d.arr[i] = i
		d.cnt[i]++
	}
	return d
}

func (d *DSU) Find(u int) int {
	if d.arr[u] != u {
		d.arr[u] = d.Find(d.arr[u])
	}
	return d.arr[u]
}

func (d *DSU) Union(a, b int) bool {
	a = d.Find(a)
	b = d.Find(b)
	if a == b {
		return false
	}
	if d.cnt[a] < d.cnt[b] {
		a, b = b, a
	}
	d.cnt[a] += d.cnt[b]
	d.arr[b] = a
	return true
}

func (d *DSU) Reset() {
	for i := 0; i < len(d.arr); i++ {
		d.arr[i] = i
		d.cnt[i] = 1
	}
}
