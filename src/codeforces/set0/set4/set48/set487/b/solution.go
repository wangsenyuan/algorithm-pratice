package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, s, l := readThreeNums(reader)
	a := readNNums(reader, n)
	res := solve(a, l, s)
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

const oo = 1 << 30

func solve(a []int, L int, s int) int {
	// 最左边的部分，至少要一个l长的部分
	// 对于任何一个i，如果它做为右边界时
	//  把i做为右边界是不行的
	// 比如j...i 不满足条件 j-1..i 可以满足条件吗？
	// 长度是没问题的，但是最大值 - 最小值呢肯定不满足条件
	// 不满足性是满足单调性的，一旦开始不满足，就一直不满足
	// 假设left[i] = j 是对于i最小的 max(a[j...i]) - min(a[j...i]) <= s 的下标
	// dp[i] = min(dp[k]) + 1 where k >= j and k < i - l
	n := len(a)

	mx := NewSegTree(n, -oo, max)
	mn := NewSegTree(n, oo, min)
	left := make([]int, n)

	findLeft := func(r int) int {
		if r == 0 {
			return r
		}
		l := left[r-1]
		for mx.Find(l, r+1)-mn.Find(l, r+1) > s {
			l++
		}

		return l
	}
	// left[i] = min index that max(left[i]..i+1) - min(left[i]...i+1) <= s

	for i := 0; i < n; i++ {
		mx.Update(i, a[i])
		mn.Update(i, a[i])
		left[i] = findLeft(i)
	}

	dp := NewSegTree(n+1, oo, min)
	dp.Update(0, 0)

	for i := 1; i <= n; i++ {
		j := left[i-1]
		if i-j >= L {
			tmp := dp.Find(j, i-L+1)
			dp.Update(i, tmp+1)
		}
	}

	ans := dp.Find(n, n+1)
	if ans >= oo {
		return -1
	}
	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, iv int, fn func(int, int) int) *SegTree {
	tr := new(SegTree)
	tr.arr = make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		tr.arr[i] = iv
	}
	tr.sz = n
	tr.initValue = iv
	tr.fn = fn
	return tr
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = tree.fn(v, tree.arr[pos])
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
