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

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		Q := make([][]int, k)
		for i := 0; i < k; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, m, E, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const B = 700

func solve(N int, M int, E [][]int, Q [][]int) []int {
	qs := make([]Query, len(Q))
	for i := 0; i < len(Q); i++ {
		l, r := Q[i][0], Q[i][1]
		qs[i] = Query{l - 1, r - 1, i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].r < qs[j].r
	})

	qv := make([][]Query, (M+B-1)/B)

	for _, q := range qs {
		i := q.l / B
		if len(qv[i]) == 0 {
			qv[i] = make([]Query, 0, 1)
		}
		qv[i] = append(qv[i], q)
	}

	var sz int
	stack := make([]Event, 0, N)

	push := func(evt Event) {
		if sz < len(stack) {
			stack[sz] = evt
			sz++
			return
		}
		// sz == len(stack)
		stack = append(stack, evt)
		sz = len(stack)
	}

	arr := make([]int, N)
	cnt := make([]int, N)

	var curTime int
	var cc int
	reset := func() {
		for i := 0; i < N; i++ {
			arr[i] = i
			cnt[i] = 1
		}
		sz = 0
		curTime = 0
		cc = N
	}

	var find func(u int) int

	find = func(u int) int {
		if arr[u] == u {
			return u
		}
		v := find(arr[u])

		push(Event{u, arr[u], cc, curTime})

		arr[u] = v

		return v
	}

	union := func(u int, v int) {
		u = find(u)
		v = find(v)
		if u == v {
			return
		}
		if cnt[u] < cnt[v] {
			u, v = v, u
		}
		// make v connect to u
		push(Event{v, arr[v], cc, curTime})

		cnt[u] += cnt[v]
		arr[v] = u
		cc--
	}

	restore := func(oldTime int) {
		for sz > 0 && oldTime < stack[sz-1].time {
			evt := stack[sz-1]
			arr[evt.id] = evt.pid
			cc = evt.cnt
			sz--
		}
		curTime = oldTime
	}

	ans := make([]int, len(Q))

	for id := 0; id < len(qv); id++ {
		reset()
		v := qv[id]
		border := (id + 1) * B
		var it int
		// v[it].r < border 那么 l..r就在同一个block中
		for ; it < len(v) && v[it].r < border; it++ {
			curTime = 1
			for i := v[it].l; i <= v[it].r; i++ {
				union(E[i][0]-1, E[i][1]-1)
			}
			ans[v[it].id] = cc
			restore(0)
		}

		jt := border
		for j := it; j < len(v); j++ {
			for jt <= v[j].r {
				union(E[jt][0]-1, E[jt][1]-1)
				jt++
			}
			curTime = 1

			for i := border - 1; i >= v[j].l; i-- {
				union(E[i][0]-1, E[i][1]-1)
			}

			ans[v[j].id] = cc

			restore(0)
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Query struct {
	l  int
	r  int
	id int
}

type Event struct {
	id   int
	pid  int
	cnt  int
	time int
}
