package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	s := readString(reader)[:n]
	ranges := make([][]int, m)
	for i := 0; i < m; i++ {
		ranges[i] = readNNums(reader, 2)
	}
	queries := make([]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNum(reader)
	}
	res := solve(s, ranges, queries)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(st string, ranges [][]int, queries []int) []int {
	n := len(st)
	s := NewSegTree(n, n+1, min)
	pos := make([]int, n)

	for i := 0; i < n; i++ {
		s.Update(i, i)
		pos[i] = -1
	}

	// v只关系st中第一次出现的位置
	var v []int

	for _, rg := range ranges {
		l, r := rg[0]-1, rg[1]-1
		for l <= r {
			l = s.Get(l, r+1)
			if l > r {
				break
			}
			pos[l] = len(v)
			v = append(v, l)
			s.Update(l, n+1)
			l++
		}
	}
	var cnt int

	for i := 0; i < n; i++ {
		if st[i] == '1' {
			cnt++
		}
	}

	var ans int

	for i := 0; i < min(cnt, len(v)); i++ {
		if st[v[i]] == '0' {
			ans++
		}
	}

	buf := []byte(st)

	res := make([]int, len(queries))

	for i, cur := range queries {
		cur--
		if pos[cur] != -1 && pos[cur] < cnt {
			if buf[cur] == '0' {
				ans--
			} else {
				ans++
			}
		}
		if buf[cur] == '0' {
			buf[cur] = '1'
			cnt++
			if cnt <= len(v) && buf[v[cnt-1]] == '0' {
				ans++
			}
		} else {
			buf[cur] = '0'
			if cnt > 0 && cnt <= len(v) && buf[v[cnt-1]] == '0' {
				ans--
			}
			cnt--
		}
		res[i] = ans
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
