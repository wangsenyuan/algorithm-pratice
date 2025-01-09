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

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d ", v))
	}
	buf.WriteByte('\n')
	os.Stdout.Write(buf.Bytes())
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
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	return res
}

func solve(a []int) []int {
	st := NewSegTree(len(a))
	n := len(a)

	for i, v := range a {
		st.update(i, pair{v, i})
	}

	res := make([]int, n)

	for {
		u := st.query(0, n)
		if u.first == inf {
			break
		}
		res[u.second] = u.first
		// 当前位要处理掉
		st.update(u.second, pair{inf, u.second})
		v := st.query(u.second+1, n)
		if v.first != u.first {
			// u.second 有可能是最后一个
			continue
		}
		res[u.second] = inf
		v.first += u.first
		st.update(v.second, v)
	}

	var ans []int
	for i := 0; i < n; i++ {
		if res[i] < inf {
			ans = append(ans, res[i])
		}
	}
	return ans
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

type SegTree struct {
	arr []pair
	sz  int
}

func min_of(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

func NewSegTree(n int) *SegTree {
	arr := make([]pair, n*2)
	for i := n; i < n*2; i++ {
		arr[i] = pair{inf, i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i*2]
	}
	return &SegTree{arr, n}
}

func (st *SegTree) update(i int, v pair) {
	i += st.sz
	st.arr[i] = v
	for i > 1 {
		st.arr[i>>1] = min_of(st.arr[i], st.arr[i^1])
		i >>= 1
	}
}

func (st *SegTree) query(l, r int) pair {
	l += st.sz
	r += st.sz
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_of(res, st.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_of(res, st.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
