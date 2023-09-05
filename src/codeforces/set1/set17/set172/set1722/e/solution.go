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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		rects := make([][]int, n)
		for i := 0; i < n; i++ {
			rects[i] = readNNums(reader, 2)
		}
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 4)
		}
		res := solve(rects, queries)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(rects [][]int, queries [][]int) []int64 {
	// w, h 只有1000， 考虑一个 dp[i][j] 表示h <= i && w <= j的总数
	dp := make([][]int64, X)
	for i := 0; i < X; i++ {
		dp[i] = make([]int64, X)
	}

	for _, cur := range rects {
		h, w := cur[0], cur[1]
		dp[h][w] += int64(h) * int64(w)
	}

	for r := 0; r < X; r++ {
		for c := 0; c < X; c++ {
			if r > 0 {
				dp[r][c] += dp[r-1][c]
			}
			if c > 0 {
				dp[r][c] += dp[r][c-1]
			}
			if r > 0 && c > 0 {
				dp[r][c] -= dp[r-1][c-1]
			}
		}
	}
	ans := make([]int64, len(queries))
	for i, cur := range queries {
		hs, ws, hb, wb := cur[0], cur[1], cur[2], cur[3]
		ans[i] = dp[hb-1][wb-1]
		ans[i] -= dp[hs][wb-1]
		ans[i] -= dp[hb-1][ws]
		ans[i] += dp[hs][ws]
	}
	return ans
}

const X = 1010

const BLOCK_SIZE = 200

func solve1(rects [][]int, queries [][]int) []int64 {
	// hs < h < hb
	// ws < w < wb
	// sum(h * w)
	// len(rects) len(queries) <= 1e5
	// h, w <= 1000
	sort.Slice(rects, func(i, j int) bool {
		return rects[i][0] < rects[j][0]
	})

	type Query struct {
		hs, ws, hb, wb int
		id             int
	}
	qs := make([]Query, len(queries))
	for i, cur := range queries {
		qs[i] = Query{cur[0], cur[1], cur[2], cur[3], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		a := qs[i].hb / BLOCK_SIZE
		b := qs[j].hb / BLOCK_SIZE
		if a != b {
			return a < b
		}
		return qs[i].hb < qs[j].hb || qs[i].hb == qs[j].hb && qs[i].hs > qs[j].hs
	})

	tr := NewSegTree(1010)

	add := func(rect []int) {
		h, w := rect[0], rect[1]
		tr.Set(w, h)
	}
	rem := func(rect []int) {
		h, w := rect[0], rect[1]
		tr.Set(w, -h)
	}

	ans := make([]int64, len(qs))
	l, r := 0, 0
	for i := 0; i < len(qs); i++ {
		for r < len(rects) && rects[r][0] < qs[i].hb {
			add(rects[r])
			r++
		}
		for r > l && rects[r-1][0] >= qs[i].hb {
			r--
			rem(rects[r])
		}
		for l > 0 && rects[l-1][0] > qs[i].hs {
			add(rects[l-1])
			l--
		}
		for l < r && rects[l][0] <= qs[i].hs {
			rem(rects[l])
			l++
		}
		ans[qs[i].id] = tr.Get(qs[i].ws+1, qs[i].wb)
	}

	return ans
}

type SegTree struct {
	sum []int64
	val []int64
	n   int
}

func NewSegTree(n int) *SegTree {
	sum := make([]int64, n)
	val := make([]int64, 2*n)
	return &SegTree{sum, val, n}
}

func (t *SegTree) Set(p int, v int) {
	t.sum[p] += int64(v)
	t.val[p+t.n] = int64(t.sum[p]) * int64(p)
	p += t.n

	for p > 1 {
		t.val[p>>1] = t.val[p] + t.val[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int64 {
	l += t.n
	r += t.n
	var res int64
	for l < r {
		if l&1 == 1 {
			res += t.val[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.val[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
