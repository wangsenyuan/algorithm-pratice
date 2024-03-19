package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	requests := readNNums(reader, n)
	res := solve(k, requests)

	var buf bytes.Buffer
	for _, row := range res {
		for i := 0; i < len(row); i++ {
			buf.WriteString(fmt.Sprintf("%d ", row[i]))
		}
		buf.WriteByte('\n')
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

func solve(k int, requests []int) [][]int {
	next := make([]*SegTree, k)
	sum := make([]*SegTree, k)

	for i := 0; i < k; i++ {
		next[i] = NewSegTree(k, min)
		for j := k; j < 2*k; j++ {
			next[i].arr[j] = j - k
		}
		for j := k - 1; j > 0; j-- {
			next[i].arr[j] = min(next[i].arr[j*2], next[i].arr[j*2+1])
		}
		// already 0
		sum[i] = NewSegTree(k, add)
	}

	center := k / 2

	calRemote := func(r int, l int, m int) int {
		nl := l + m - 1
		ans := abs(r-center) * m
		if nl <= center {
			// in the first half
			ans += m*center - countSum(l, nl)
		} else if center <= l {
			ans += countSum(l, nl) - m*center
		} else {
			x := center - l
			ans += x*center - countSum(l, l+x-1)
			ans += countSum(center, nl) - (m-x)*center
		}

		return ans
	}

	check := func(r int, x int) (ok bool, remote int, l int) {
		remote = -1
		l = -1

		for i := 0; i+x <= k; {
			cnt := sum[r].Get(i, i+x, 0)
			if cnt == 0 {
				tmp := calRemote(r, i, x)
				if remote < 0 || tmp < remote {
					remote = tmp
					l = i
				}
				if tmp == 0 {
					return true, remote, i
				}
				// 至少要移动一个位置
				cnt++
			}
			// 假设cnt个都在最左边（理想情况)下一个地方肯定在l+cnt 开始
			i = next[r].Get(i+cnt, k, k)
		}

		if remote >= 0 {
			return remote == 0, remote, l
		}

		return false, -1, -1
	}

	take := func(r int, l int, m int) {
		for i := l; i < l+m; i++ {
			// take up the position
			next[r].Update(i, k)
			sum[r].Update(i, 1)
		}
	}

	process := func(x int) []int {
		ans := []int{-1, -1, -1}
		for d := 0; center-d >= 0; d++ {
			ok1, v1, l1 := check(center-d, x)
			ok2, v2, l2 := false, -1, -1
			if d != 0 {
				ok2, v2, l2 = check(center+d, x)
			}
			if v1 != -1 && (ans[1] < 0 || v1 <= ans[1]) {
				ans[0] = center - d
				ans[1] = v1
				ans[2] = l1
			}
			if v2 != -1 && (ans[1] < 0 || v2 < ans[1]) {
				ans[0] = center + d
				ans[1] = v2
				ans[2] = l2
			}
			// 有一行是空的
			if ok1 || ok2 {
				break
			}
		}
		if ans[0] < 0 {
			return []int{-1}
		}
		take(ans[0], ans[2], x)

		return []int{ans[0] + 1, ans[2] + 1, ans[2] + x}
	}

	ans := make([][]int, len(requests))
	for i, x := range requests {
		ans[i] = process(x)
	}

	return ans
}

func countSum(i int, j int) int {
	// i, i + 1, i + 2, ... j
	return (i + j) * (j - i + 1) / 2
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func add(a, b int) int {
	return a + b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type SegTree struct {
	arr []int
	sz  int
	f   func(int, int) int
}

func NewSegTree(n int, f func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n, f}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = t.f(t.arr[p], v)
	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int, iv int) int {
	res := iv
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
