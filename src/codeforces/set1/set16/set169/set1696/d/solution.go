package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		arr := readNNums(reader, n)
		res := solve(arr)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(arr []int) int {
	// a[i] -> a[j] if max(a[i..j]), min(a[i..j]) = a[i], a[j]
	// 1， 2， 5， 3， 4
	// 是不是连续的单调的部分，可以直接处理掉的？
	n := len(arr)

	ids := make([]Pair, n)
	for i := 0; i < n; i++ {
		ids[i] = Pair{arr[i], i}
	}

	mn := NewSegTree(ids, minPair)
	mx := NewSegTree(ids, maxPair)

	var dfs func(l int, r int) int

	dfs = func(l int, r int) int {
		if l == r {
			return 0
		}

		x := mx.Get(l, r+1, Pair{arr[r], r})
		// x is the critical point, will always be used
		i := x.second

		var ans int
		if l < i {
			y := mn.Get(l, i, Pair{arr[l], l})
			j := y.second
			ans = 1 + dfs(l, j)
		}
		if i < r {
			y := mn.Get(i, r+1, Pair{arr[i], i})
			j := y.second
			ans += 1 + dfs(j, r)
		}
		return ans
	}

	return dfs(0, n-1)
}

func solve1(arr []int) int {
	// a[i] -> a[j] if max(a[i..j]), min(a[i..j]) = a[i], a[j]
	// 1， 2， 5， 3， 4
	// 是不是连续的单调的部分，可以直接处理掉的？
	n := len(arr)
	x := make([]int, n)
	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		x[i] = n
		for p > 0 && arr[stack[p-1]] < arr[i] {
			x[stack[p-1]] = i
			p--
		}
		stack[p] = i
		p++
	}

	p = 0
	y := make([]int, n)
	for i := 0; i < n; i++ {
		y[i] = n
		for p > 0 && arr[stack[p-1]] > arr[i] {
			y[stack[p-1]] = i
			p--
		}
		stack[p] = i
		p++
	}

	ids := make([]Pair, n)
	for i := 0; i < n; i++ {
		ids[i] = Pair{arr[i], i}
	}

	mn := NewSegTree(ids, minPair)
	mx := NewSegTree(ids, maxPair)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = n
	}
	dp[0] = 0
	for i := 0; i+1 < n; i++ {
		if arr[i+1] > arr[i] {
			//只能找最大值
			tmp := mx.Get(i+1, y[i], Pair{arr[i+1], i + 1})
			j := tmp.second
			dp[j] = min(dp[j], dp[i]+1)
		} else {
			tmp := mn.Get(i+1, x[i], Pair{arr[i+1], i + 1})
			j := tmp.second
			dp[j] = min(dp[j], dp[i]+1)
		}
		dp[i+1] = min(dp[i+1], dp[i]+1)
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

func maxPair(a, b Pair) Pair {
	if a.first > b.first {
		return a
	}
	return b
}

func minPair(a, b Pair) Pair {
	if a.first < b.first {
		return a
	}
	return b
}

type SegTree struct {
	arr []Pair
	n   int
	fn  func(Pair, Pair) Pair
}

func NewSegTree(in []Pair, fn func(Pair, Pair) Pair) *SegTree {
	tr := new(SegTree)
	tr.n = len(in)
	tr.arr = make([]Pair, tr.n*2)
	copy(tr.arr[tr.n:], in)
	for i := tr.n - 1; i > 0; i-- {
		tr.arr[i] = fn(tr.arr[2*i], tr.arr[i*2+1])
	}
	tr.fn = fn
	return tr
}

func (tr *SegTree) Get(l int, r int, res Pair) Pair {
	l += tr.n
	r += tr.n
	for l < r {
		if l&1 == 1 {
			res = tr.fn(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = tr.fn(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
