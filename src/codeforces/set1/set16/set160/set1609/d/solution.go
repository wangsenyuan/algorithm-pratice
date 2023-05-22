package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, d := readTwoNums(reader)
		cond := make([][]int, d)
		for i := 0; i < d; i++ {
			cond[i] = readNNums(reader, 2)
		}
		res := solve(n, cond)
		for i := 0; i < d; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(n int, cond [][]int) []int {
	// 满足前i个条件的情况下，进行i次介绍。但是这个介绍不和要求一次
	// 假设 (1, 2) (2, 3), 现在 要求(1, 3)认识，在他们已经认识的情况下，其实可以再进行一次介绍，比如 (1, 4)
	// 所以，这里需要知道前面进行了几次有效的介绍
	// 然后剩余的次数，应该是去连接这些分组，在分组连接完成的情况下，在添加新的节点进去
	// 有效的介绍很简单
	set := NewDSU(n)

	arr := make([]Pair, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = Pair{1, i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = max_pair(arr[i*2], arr[2*i+1])
	}

	update := func(p int, v int) {
		p += n
		arr[p] = arr[p].UpdateFirst(v)
		for p > 1 {
			arr[p>>1] = max_pair(arr[p], arr[p^1])
			p >>= 1
		}
	}

	query := func(l int, r int) Pair {
		l += n
		r += n
		res := Pair{0, 0}
		for l < r {
			if l&1 == 1 {
				res = max_pair(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max_pair(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	ans := make([]int, len(cond))
	// n <= 1000
	pending := make([]Pair, n)
	var avi int
	for i, cd := range cond {
		// can use pu & pv to merge the largest two
		if i > 0 {
			ans[i] = ans[i-1]
		}
		if ans[i] == n-1 {
			continue
		}

		u, v := cd[0]-1, cd[1]-1
		pu, pv := set.Find(u), set.Find(v)
		if pu != pv {
			update(pu, 0)
			update(pv, 0)
			set.Union(u, v)
			pu = set.Find(u)
			update(pu, set.cnt[pu])

			ans[i] = set.cnt[pu] - 1

			if i > 0 {
				ans[i] = max(ans[i], ans[i-1])
			}
		} else {
			avi++
		}

		tmp := avi
		var j int
		var sum int
		for tmp >= 0 && sum < n {
			cur := query(0, n)
			sum += cur.first
			pending[j] = cur
			j++
			update(cur.second, 0)
			tmp--
		}

		ans[i] = max(ans[i], sum-1)
		for j > 0 {
			cur := pending[j-1]
			update(cur.second, cur.first)
			j--
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func max_pair(a, b Pair) Pair {
	if a.first > b.first {
		return a
	}
	return b
}

func (this Pair) UpdateFirst(v int) Pair {
	return Pair{v, this.second}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type DSU struct {
	arr  []int
	cnt  []int
	find func(int) int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}

	var find func(x int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	return &DSU{arr, cnt, find}
}

func (dsu *DSU) Find(x int) int {
	return dsu.find(x)
}

func (dsu *DSU) Union(x, y int) bool {
	x = dsu.find(x)
	y = dsu.find(y)
	if x == y {
		return false
	}
	if dsu.cnt[x] < dsu.cnt[y] {
		x, y = y, x
	}
	dsu.cnt[x] += dsu.cnt[y]
	dsu.arr[y] = x
	return true
}
