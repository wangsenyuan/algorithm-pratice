package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readNum(reader)
	s := readString(reader)
	res := solve(s)
	s = fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func solve(s string) []int {
	// +1 for left, -1 for right
	n := len(s)

	tr := Build(n)

	var pos int

	buf := make([]int, n)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			pos = max(pos-1, 0)
		} else if s[i] == 'R' {
			pos++
		} else {
			v := get(s[i])
			w := buf[pos]
			buf[pos] = v
			v -= w
			tr.Update(pos, n-1, v)
		}
		val := tr.Get(0, n-1)
		last := tr.Get(n-1, n-1)
		if val.first != 0 || last.first != 0 {
			ans[i] = -1
		} else {
			ans[i] = val.second
		}
	}

	return ans
}

func get(x byte) int {
	if x == '(' {
		return 1
	}
	if x == ')' {
		return -1
	}
	return 0
}

func bruteForce(s string) []int {
	n := len(s)

	buf := make([]int, n)

	check := func() int {
		var level int
		var res int
		for i := 0; i < n; i++ {
			level += buf[i]
			res = max(res, level)
			if level < 0 {
				return -1
			}
		}
		if level != 0 {
			return -1
		}
		return res
	}

	ans := make([]int, n)
	var pos int
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			pos = max(pos-1, 0)
		} else if s[i] == 'R' {
			pos++
		} else {
			v := get(s[i])
			buf[pos] = v
		}
		ans[i] = check()
	}

	return ans
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

type SegTree struct {
	mn   []int
	mx   []int
	lazy []int
	sz   int
}

func Build(n int) *SegTree {
	mn := make([]int, 4*n)
	mx := make([]int, 4*n)
	lazy := make([]int, 4*n)

	return &SegTree{mn, mx, lazy, n}
}

func (tr *SegTree) update(i int, v int) {
	tr.mn[i] += v
	tr.mx[i] += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(i*2+1, tr.lazy[i])
		tr.update(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) pull(i int) {
	tr.mn[i] = min(tr.mn[i*2+1], tr.mn[i*2+2])
	tr.mx[i] = max(tr.mx[i*2+1], tr.mx[i*2+2])
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		tr.pull(i)
	}

	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) pair {
	var loop func(i int, l int, r int) pair
	loop = func(i int, l int, r int) pair {
		if R < l || r < L {
			return pair{inf, -inf}
		}

		if L <= l && r <= R {
			return pair{tr.mn[i], tr.mx[i]}
		}

		tr.push(i)

		mid := (l + r) / 2

		a := loop(i*2+1, l, mid)
		b := loop(i*2+2, mid+1, r)

		return pair{min(a.first, b.first), max(a.second, b.second)}
	}

	return loop(0, 0, tr.sz-1)
}
