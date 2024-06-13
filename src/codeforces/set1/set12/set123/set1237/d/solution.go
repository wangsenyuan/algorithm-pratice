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

	n := readNum(reader)
	a := readNNums(reader, n)

	var buf bytes.Buffer

	res := solve(a)

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

const inf = 1 << 30

func solve(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = inf
	}
	que := make([]Pair, 6*n)
	var front, tail int
	for i, j := 0, 0; i < n; i++ {
		for j < 3*n && (tail == front || a[j%n]*2 >= que[tail].first) {
			for front > tail && que[front-1].first <= a[j%n] {
				front--
			}
			que[front] = Pair{a[j%n], j}
			front++
			j++
		}

		if j < 3*n {
			ans[i] = j - i
		}

		if tail < front && que[tail].second == i {
			tail++
		}
	}

	for i := 0; i < n; i++ {
		if ans[i] == inf {
			ans[i] = -1
		}
	}
	return ans
}

func solve2(a []int) []int {
	n := len(a)

	// stack 从top到bot降序排列
	stack := make([]int, 2*n)
	var top int

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = inf
	}

	for i := 2*n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]%n] >= a[i%n] {
			top--
		}

		if top > 0 {
			l, r := 0, top
			for l < r {
				mid := (l + r) / 2
				if a[stack[mid]%n]*2 < a[i%n] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			r--
			// a[stack[r]] * 2 < a[i % n]
			if r >= 0 {
				ans[i%n] = min(stack[r]-i, ans[i%n])
			}
		}
		stack[top] = i
		top++
	}

	for i := 2*n - 1; i >= 0; i-- {
		ans[i%n] = min(ans[i%n], ans[(i+1)%n]+1)
	}
	for i := 0; i < n; i++ {
		if ans[i] == inf {
			ans[i] = -1
		}
	}
	return ans
}

func solve1(a []int) []int {
	n := len(a)

	tr := Build(n)
	for i := 0; i < n; i++ {
		tr.Set(i, a[i])
	}

	arr := make([]Pair, n)

	for i := 0; i < n; i++ {
		arr[i] = Pair{a[i], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first > arr[j].first || arr[i].first == arr[j].first && arr[i].second > arr[j].second
	})

	sg := NewSegTree(n)

	ans := make([]int, n)

	for _, pt := range arr {
		val := pt.first/2 - (1 - pt.first%2)
		i := pt.second
		ans[i] = -1
		j := tr.FindLeftMostExpectPosition(val, i, n-1)
		if j == inf {
			j = tr.FindLeftMostExpectPosition(val, 0, i)
		}
		if j == inf {
			j = sg.Get(i, n)
			if j == inf {
				j = sg.Get(0, i)
			}
			if j < inf && ans[j] > 0 {
				ans[i] = (j+n-i)%n + ans[j]
			}
		} else {
			ans[i] = (j + n - i) % n
			if j > i {
				k := sg.Get(i, j)
				if k < inf && ans[k] > 0 {
					ans[i] = k - i + ans[k]
				}
			} else {
				k := sg.Get(i, n)
				if k == inf {
					k = sg.Get(0, j)
				}
				if k < inf && ans[k] > 0 {
					ans[i] = (k+n-i)%n + ans[k]
				}
			}
		}
		sg.update(i, i)
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

type Tree struct {
	val []int
	sz  int
}

func Build(n int) *Tree {
	arr := make([]int, 4*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}

	return &Tree{arr, n}
}

func (tr *Tree) Set(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.val[i] = min(tr.val[i*2+1], tr.val[i*2+2])
	}

	loop(0, 0, tr.sz-1)
}

func (tr *Tree) FindLeftMostExpectPosition(expect_val int, L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if R < l || r < L || tr.val[i] > expect_val {
			return inf
		}
		mid := (l + r) / 2
		if L <= l && r <= R {
			if l == r {
				return l
			}
			// l < r
			if tr.val[2*i+1] <= expect_val {
				return loop(i*2+1, l, mid)
			}
			return loop(i*2+2, mid+1, r)
		}
		return min(loop(2*i+1, l, mid), loop(2*i+2, mid+1, r))
	}

	return loop(0, 0, tr.sz-1)
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) SegTree {
	tree := make([]int, 2*n)

	for i := 0; i < n; i++ {
		tree[i+n] = inf
	}

	for i := n - 1; i > 0; i-- {
		tree[i] = min(tree[i<<1], tree[i<<1|1])
	}

	return SegTree{tree, n}
}

func (tree *SegTree) update(p int, v int) {
	p += tree.n
	tree.arr[p] = v

	for p > 0 {
		tree.arr[p>>1] = min(tree.arr[p], tree.arr[p^1])
		p >>= 1
	}
}

func (tree SegTree) Get(l int, r int) int {
	l += tree.n
	r += tree.n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
