package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	v := readNNums(reader, n)

	res := solve(k, v)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(k int, v []int) int {
	sort.Ints(v)
	n := len(v)

	dp := make([]int, n+1)

	far := make([]int, n+1)
	check := func(expect int) bool {
		for i, r := 0, 0; i < n; i++ {
			for r < n && v[r]-v[i] <= expect {
				r++
			}
			far[i] = r
			dp[i] = 0
		}
		var sum int

		for i, r := n-k, n; i >= 0; i-- {
			dp[i] = 0
			for r > far[i] {
				sum -= dp[r]
				r--
			}
			sum += dp[i+k]
			if far[i] == n || i+k <= far[i] && sum > 0 {
				dp[i] = 1
			}
		}

		return dp[0] == 1
	}

	l, r := 0, v[n-1]-v[0]

	if k == n {
		return r
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func solve1(k int, v []int) int {
	sort.Ints(v)
	n := len(v)

	// 假设区间 [l...r]的size < k
	// 也就是说 v[r+1] - v[l] > expect
	// 这时有几种可能性
	// 1. 不存在答案
	// 2 将区间分成两部分，前半部分和前面的区间合并，后半部分和后面的区间合并
	// 3.拆分区间，会影响到其他的区间分配
	// 所以这里肯定有什么东西被我忽略了
	// 对于每个i，是可以知道从它这里开始（或结束）是否能把全部按照条件分配好
	// dp[i] = 在区间 [i+k...far(i)] 中间存在一个合理的划分
	far := make([]int, n+1)
	dp := NewSegTree(n, 0, plus)
	check := func(expect int) bool {
		for i, r := 0, 0; i < n; i++ {
			for r < n && v[r]-v[i] <= expect {
				r++
			}
			far[i] = r
			dp.Update(i, 0)
		}

		for i := n - k; i >= 0; i-- {
			if far[i] == n {
				dp.Update(i, 1)
			} else if i+k <= far[i] {
				cnt := dp.Get(i+k, far[i]+1)
				if cnt > 0 {
					dp.Update(i, 1)
				}
			}
		}

		return dp.Get(0, 1) == 1
	}

	l, r := 0, v[n-1]-v[0]

	if k == n {
		return r
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func plus(a, b int) int {
	return a + b
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
