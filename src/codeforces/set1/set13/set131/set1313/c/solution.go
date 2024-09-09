package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(a []int) []int {
	n := len(a)

	stack := make([]int, n)
	var top int

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		for top > 0 && a[stack[top-1]] > a[i] {
			top--
		}
		if top == 0 {
			dp[i] = (i + 1) * a[i]
		} else {
			j := stack[top-1]
			dp[i] = dp[j] + (i-j)*a[i]
		}
		stack[top] = i
		top++
	}

	best := dp[n-1]
	at := n - 1

	fp := make([]int, n)
	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] > a[i] {
			top--
		}
		if top == 0 {
			fp[i] = (n - i) * a[i]
		} else {
			j := stack[top-1]
			fp[i] = fp[j] + (j-i)*a[i]
		}

		if dp[i]+fp[i]-a[i] > best {
			best = dp[i] + fp[i] - a[i]
			at = i
		}

		stack[top] = i
		top++
	}

	res := make([]int, n)
	res[at] = a[at]
	for i := at - 1; i >= 0; i-- {
		res[i] = min(a[i], res[i+1])
	}

	for i := at + 1; i < n; i++ {
		res[i] = min(a[i], res[i-1])
	}

	return res
}

func solve1(a []int) []int {
	n := len(a)

	type pair struct {
		first  int
		second int
	}

	nums := make([]pair, n)
	for i := 0; i < n; i++ {
		nums[i] = pair{a[i], i}
	}

	slices.SortFunc(nums, func(a, b pair) int {
		if a.first != b.first {
			return a.first - b.first
		}
		return a.second - b.second
	})

	set := NewSegTree(n, -1, max)

	dp := make([]int, n)

	for _, cur := range nums {
		i := cur.second
		j := set.Get(0, i)
		if j < 0 {
			// a[i]是最小值
			dp[i] = (i + 1) * a[i]
		} else {
			// 这个应该是对的
			dp[i] = (i-j)*a[i] + dp[j]
		}
		set.Update(i, i)
	}

	set = NewSegTree(n, n, min)

	slices.SortFunc(nums, func(a, b pair) int {
		if a.first != b.first {
			return a.first - b.first
		}
		// 下标大的排在后面
		return b.second - a.second
	})

	best := dp[n-1]
	at := n - 1
	fp := make([]int, n)

	for _, cur := range nums {
		i := cur.second
		j := set.Get(i+1, n)
		if j == n {
			fp[i] = (n - i) * a[i]
		} else {
			fp[i] = (j-i)*a[i] + fp[j]
		}
		if dp[i]+fp[i]-a[i] > best {
			best = dp[i] + fp[i] - a[i]
			at = i
		}
		set.Update(i, i)
	}

	res := make([]int, n)
	res[at] = a[at]
	for i := at - 1; i >= 0; i-- {
		res[i] = min(a[i], res[i+1])
	}

	for i := at + 1; i < n; i++ {
		res[i] = min(a[i], res[i-1])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, iv int, fn func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		arr[i] = iv
	}
	return &SegTree{arr, n, iv, fn}
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Get(l, r int) int {
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
