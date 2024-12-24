package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	x, s := make([]int, n), make([]int, n)
	for i := range n {
		x[i], s[i] = readTwoNums(reader)
	}
	return solve(x, s, m)
}

const inf = 1 << 60

func solve(x []int, s []int, m int) int {
	n := len(x)
	// n <= 80
	type pair struct {
		first  int
		second int
	}

	ants := make([]pair, n)
	for i := range n {
		ants[i] = pair{x[i], s[i]}
	}

	slices.SortFunc(ants, func(a, b pair) int {
		return a.first - b.first
	})

	tr1 := NewSegTree(m + 1)
	tr1.Update(0, 0)
	tr2 := NewSegTree(m + 1)
	tr2.Update(0, 0)
	dp := make([]int, m+1)

	for _, cur := range ants {
		x, s := cur.first, cur.second

		for i := 1; i <= m; i++ {
			dp[i] = inf
			// dp[i] = ?
			// 如果前面的cover到的位置是[1...y]
			// 如果 y <= i < x
			// 且 x - y >= s => y <= x - s, 这时只需要考虑把ycover住
			// 需要增加的高度 = dp[y] - y + x - s

			if i < x-s {
				// 如果处在当前塔的范围外面
				dp[i] = min(dp[i], tr2.Get(0, i)+x-s-1)
			} else if i <= x+s {
				// 如果处在当前塔的范围内
				dp[i] = min(dp[i], tr1.Get(max(0, x-s), i))
				dp[i] = min(dp[i], tr2.Get(0, max(0, x-s))+x-s-1)
			} else {
				// 如果存在当前塔的范围后面
				// 这个貌似比较麻烦
				j := x - (i - x)
				// 如果在j的前面, 那么以j为主
				if j >= 1 {
					dp[i] = min(dp[i], tr2.Get(0, j)+x-s-1)
				}
				// 在j的后面以i为主
				dp[i] = min(dp[i], tr1.Get(j, i)+i-x-s)
			}
		}
		for i := 1; i <= m; i++ {
			tr1.Update(i, dp[i])
			tr2.Update(i, dp[i]-i)
		}
	}

	return tr1.Get(m, m+1)
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := range 2 * n {
		arr[i] = inf
	}
	return &SegTree{arr, n}
}

func (tr *SegTree) Update(p int, v int) {
	p += tr.sz
	tr.arr[p] = min(tr.arr[p], v)
	for p > 1 {
		tr.arr[p>>1] = min(tr.arr[p>>1], v)
		p >>= 1
	}
}

func (tr *SegTree) Get(l int, r int) int {
	res := inf
	l += tr.sz
	r += tr.sz
	for l < r {
		if l&1 == 1 {
			res = min(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
