package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	s := readString(reader)

	res := solve(n, k, s)

	fmt.Println(res)
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

const inf = 1 << 60

func solve(n int, k int, s string) int {
	set := NewSegTree(n + 1)

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			set.Update(i+1, i+1)
		}
	}

	dp := make([]int, n+2)
	for i := 0; i <= n; i++ {
		dp[i] = inf
	}

	for i := n; i > 0; i-- {
		dp[i] = min(dp[i], dp[i+1]+i)

		j := set.Get(max(i-k, 1), min(i+k, n+1))
		if j < inf {
			ii := max(1, j-k)
			dp[ii] = min(dp[ii], dp[i+1]+j)
		}
	}

	return dp[1]
}

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = min(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := inf
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = min(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func solve1(n int, k int, s string) int {
	next := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		next[i] = inf
	}
	pos := 2 * n
	for i := n - 1; i >= -k; i-- {
		if i >= 0 && s[i] == '1' {
			pos = i
		}
		if i+k < n {
			next[i+k] = pos
		}
	}
	dp := make([]int, n+2)
	for i := 0; i <= n; i++ {
		dp[i] = inf
	}
	dp[n] = 0
	for i := n; i > 0; i-- {
		if next[i-1]-(i-1) <= k {
			nw := max(0, next[i-1]-k)
			dp[nw] = min(dp[nw], dp[i]+next[i-1]+1)
		}
		dp[i-1] = min(dp[i-1], dp[i]+i)
	}

	return dp[0]
}
