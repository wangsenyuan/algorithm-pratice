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
		n := readNum(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(a []int, b []int) bool {
	n := len(a)
	// r表示当前选择的最大的区间
	mx := NewOpSet(n, 0, max)

	for i := 0; i < n; i++ {
		mx.Update(i, a[i])
	}

	pos := make(map[int]int)
	L := make([]int, n)

	for i := 0; i < n; i++ {
		pos[a[i]] = i
		if a[i] > b[i] {
			return false
		}
		if j, ok := pos[b[i]]; ok && mx.Get(j, i, 0) <= b[i] {
			L[i] = j
		} else {
			L[i] = -1
		}
	}
	pos = make(map[int]int)
	R := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		pos[a[i]] = i
		if j, ok := pos[b[i]]; ok && mx.Get(i, j+1, 0) <= b[i] {
			R[i] = j
		} else {
			R[i] = n
		}
	}

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return b[id[i]] < b[id[j]]
	})

	vis := NewOpSet(n, 0, add)

	for i := 0; i < n; {
		j := i
		for i < n && b[id[i]] == b[id[j]] {
			l, r := L[id[i]], R[id[i]]

			ok := l >= 0 && vis.Get(l, id[i], 0) == 0 || r < n && vis.Get(id[i], r, 0) == 0

			if !ok {
				return false
			}

			i++
		}
		// delay update
		for j < i {
			vis.Update(id[j], 1)
			j++
		}
	}

	return true
}

func add(a, b int) int {
	return a + b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type OpSet struct {
	arr []int
	sz  int
	op  func(int, int) int
}

func NewOpSet(n int, iv int, op func(int, int) int) *OpSet {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &OpSet{arr, n, op}
}

func (set *OpSet) Update(p int, v int) {
	p += set.sz
	set.arr[p] = set.op(set.arr[p], v)
	for p > 1 {
		set.arr[p>>1] = set.op(set.arr[p], set.arr[p^1])
		p >>= 1
	}
}

func (set *OpSet) Get(l int, r int, iv int) int {
	l += set.sz
	r += set.sz
	res := iv
	for l < r {
		if l&1 == 1 {
			res = set.op(res, set.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = set.op(res, set.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
