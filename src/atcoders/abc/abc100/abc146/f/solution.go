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
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) []int {
	n, k := readTwoNums(reader)
	s := readString(reader)[:n+1]
	return solve(k, s)
}

func solve(m int, s string) []int {
	n := len(s)

	dp := NewSegTree(n)
	dp.Update(n-1, 0)
	next := make([]int, n)
	next[n-1] = n
	for i := n - 2; i >= 0; i-- {
		next[i] = n
		if s[i] == '1' {
			continue
		}

		tmp := dp.Get(i+1, min(n, i+m+1))
		if tmp.first == inf {
			return nil
		}
		next[i] = tmp.second
		dp.Update(i, tmp.first+1)
	}
	var ans []int
	for i := 0; i < n-1; i = next[i] {
		ans = append(ans, next[i]-i)
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type pair struct {
	first  int
	second int
}

func min_pair(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

const inf = 1 << 60

type SegTree []pair

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{inf, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i*2], arr[i*2+1])
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p].first = v

	for p > 1 {
		tr[p>>1] = min_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_pair(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_pair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
