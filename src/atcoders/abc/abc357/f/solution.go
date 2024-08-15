package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')

		if s[0] == '3' {
			queries[i] = make([]int, 3)
		} else {
			queries[i] = make([]int, 4)
		}
		var pos int
		for j := 0; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}

	res := solve(a, b, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(a []int, b []int, queries [][]int) []int {
	n := len(a)

	arr := make([]data, n)
	for i := 0; i < n; i++ {
		arr[i] = data{a[i] % mod, b[i] % mod, mul(a[i], b[i])}
	}

	t := make(seg, 4*n)
	t.build(arr, 1, 1, n)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			l, r, x := cur[1], cur[2], cur[3]
			t.update(1, l, r, pair{x, 0})
		} else if cur[0] == 2 {
			l, r, x := cur[1], cur[2], cur[3]
			t.update(1, l, r, pair{0, x})
		} else {
			l, r := cur[1], cur[2]
			ans = append(ans, t.query(1, l, r).sab)
		}
	}

	return ans
}

type data struct{ sa, sb, sab int }
type pair struct{ a, b int }
type seg []struct {
	l, r int
	d    data
	todo pair
}

var todoInit pair

// 线段树模板，只需要实现 mergeInfo 和 do，其余都是固定的
func mergeInfo(a, b data) data {
	return data{
		(a.sa + b.sa) % mod,
		(a.sb + b.sb) % mod,
		(a.sab + b.sab) % mod,
	}
}

func (t seg) do(O int, p pair) {
	o := &t[O]

	sz := o.r - o.l + 1
	o.d.sa = (o.d.sa + sz*p.a) % mod
	o.d.sab = (o.d.sab + o.d.sb*p.a) % mod
	o.d.sb = (o.d.sb + sz*p.b) % mod
	o.d.sab = (o.d.sab + o.d.sa*p.b) % mod

	o.todo.a = (o.todo.a + p.a) % mod
	o.todo.b = (o.todo.b + p.b) % mod
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != todoInit {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit
	}
}

func (t seg) build(a []data, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].d = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r int, v pair) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func (t seg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].d
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}
