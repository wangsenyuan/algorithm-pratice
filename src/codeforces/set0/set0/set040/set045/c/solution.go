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
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) [][]int {
	n := readNum(reader)
	s := readString(reader)
	a := readNNums(reader, n)
	return solve(s, a)
}

func solve(s string, a []int) [][]int {
	n := len(s)
	next := make([]int, n)
	prev := make([]int, n)
	tr := NewSegTree(n)
	for i := 0; i < n; i++ {
		prev[i] = i - 1
		next[i] = i + 1
		if i+1 < n && s[i] != s[i+1] {
			tr.Update(i, abs(a[i+1]-a[i]))
		} else {
			tr.Update(i, inf)
		}
	}

	var res [][]int
	for {
		tmp := tr.Get(0, n)
		if tmp.first == inf {
			break
		}
		i := tmp.second
		j := next[i]
		res = append(res, []int{i + 1, j + 1})
		tr.Update(i, inf)
		tr.Update(j, inf)
		i = prev[i]
		j = next[j]
		if i >= 0 && j < n {
			if s[i] != s[j] {
				tr.Update(i, abs(a[j]-a[i]))
			} else {
				tr.Update(i, inf)
			}
			next[i] = j
			prev[j] = i
		} else if i >= 0 {
			tr.Update(i, inf)
			next[i] = n
		} else if j < n {
			prev[j] = -1
		}
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func min_of(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

type SegTree struct {
	arr []pair
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = pair{inf, i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = min_of(arr[i*2], arr[i*2+1])
	}
	return &SegTree{arr, n}
}

func (tr *SegTree) Update(p int, v int) {
	p += tr.sz
	tr.arr[p].first = v
	for p > 0 {
		tr.arr[p>>1] = min_of(tr.arr[p], tr.arr[p^1])
		p >>= 1
	}
}

func (tr *SegTree) Get(l int, r int) pair {
	l += tr.sz
	r += tr.sz
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_of(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_of(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
