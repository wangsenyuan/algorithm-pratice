package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	users := make([][]int, n)
	for i := 0; i < n; i++ {
		users[i] = readNNums(reader, 2)
	}
	return solve(users)
}

func solve(users [][]int) []int {

	pos := make(map[int]int)
	for _, user := range users {
		pos[user[0]]++
		pos[user[1]]++
	}
	var arr []int
	for k := range pos {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	m := len(arr)
	for i := range m {
		pos[arr[i]] = i
	}

	open := make([][]int, m)
	end := make([][]int, m)

	for i, user := range users {
		l, r := pos[user[0]], pos[user[1]]
		open[l] = append(open[l], i)
		end[r] = append(end[r], i)
	}
	n := len(users)

	ans := make([]int, n)

	mn := NewSegTree(m, inf, func(a, b int) int {
		return min(a, b)
	})

	cnt := make([]int, m)

	type pair struct {
		first  int
		second int
	}

	for i := 0; i < m; i++ {
		var buf []pair
		for _, j := range open[i] {
			r := pos[users[j][1]]
			cnt[r]++
			buf = append(buf, pair{r, j})
		}

		slices.SortFunc(buf, func(a, b pair) int {
			return b.first - a.first
		})

		for _, cur := range buf {
			j := cur.second
			r := pos[users[j][1]]

			if cnt[r] == 1 {
				// 找到r后面的最小值
				y := mn.Find(r, m)
				if y < inf {
					ans[j] += y - users[j][1]
				}
			}
			mn.Update(r, users[j][1])
		}
	}

	mx := NewSegTree(m, -1, func(a, b int) int {
		return max(a, b)
	})

	clear(cnt)

	for i := m - 1; i >= 0; i-- {
		var buf []pair
		for _, j := range end[i] {
			l := pos[users[j][0]]
			cnt[l]++
			buf = append(buf, pair{l, j})
		}

		slices.SortFunc(buf, func(a, b pair) int {
			return a.first - b.first
		})
		for _, cur := range buf {
			j := cur.second
			l := pos[users[j][0]]
			if cnt[l] == 1 {
				x := mx.Find(0, l)
				if x >= 0 {
					ans[j] += users[j][0] - x
				}
			}
			mx.Update(l, users[j][0])
		}
	}

	return ans
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

func solve1(users [][]int) []int {
	n := len(users)

	ans := make([]int, n)

	type person struct {
		id int
		l  int
		r  int
	}

	people := make([]person, n)
	for i, user := range users {
		people[i] = person{i, user[0], user[1]}
	}

	slices.SortFunc(people, func(a, b person) int {
		if a.l != b.l {
			return a.l - b.l
		}
		return b.r - a.r
	})

	var tr *Tree

	for i := 0; i < n; {
		j := i

		for i < n && people[i].l == people[j].l {
			tr = tr.Set(people[i].r, 0, 1e9, 1)
			i++
		}
		for j < i {
			cur := people[j]
			id := cur.id
			tr = tr.Set(cur.r, 0, 1e9, -1)
			r := tr.GetValue1(0, 1e9, cur.r, 1e9)
			if r <= 1e9 {
				ans[id] += r - cur.r
			}
			tr = tr.Set(cur.r, 0, 1e9, 1)
			j++
		}
	}

	var tr2 *Tree

	slices.SortFunc(people, func(a, b person) int {
		if a.r != b.r {
			return b.r - a.r
		}
		return a.l - b.l
	})

	for i := 0; i < n; {
		j := i
		for i < n && people[i].r == people[j].r {
			tr2 = tr2.Set(people[i].l, 0, 1e9, 1)
			i++
		}
		for j < i {
			cur := people[j]
			id := cur.id
			tr2 = tr2.Set(cur.l, 0, 1e9, -1)
			l := tr2.GetValue2(0, 1e9, 0, cur.l)
			if l >= 0 {
				ans[id] += cur.l - l
			}
			tr2 = tr2.Set(cur.l, 0, 1e9, 1)
			j++
		}
	}

	return ans
}

const inf = 1 << 30

type Tree struct {
	left, right *Tree
	val1        int
	val2        int
	cnt         int
}

func (tr *Tree) Value1() int {
	if tr == nil {
		return inf
	}
	return tr.val1
}

func (tr *Tree) Value2() int {
	if tr == nil {
		return -inf
	}
	return tr.val2
}

func (tr *Tree) Set(p int, l int, r int, v int) *Tree {
	if tr == nil {
		tr = new(Tree)
		tr.val1 = inf
		tr.val2 = -inf
		tr.cnt = 0
	}
	if l == r {
		tr.cnt += v
		if tr.cnt > 0 {
			tr.val1 = l
			tr.val2 = l
		} else {
			tr.val1 = inf
			tr.val2 = -inf
		}

		return tr
	}
	mid := (l + r) / 2
	if p <= mid {
		tr.left = tr.left.Set(p, l, mid, v)
	} else {
		tr.right = tr.right.Set(p, mid+1, r, v)
	}
	tr.val1 = min(tr.left.Value1(), tr.right.Value1())
	tr.val2 = max(tr.left.Value2(), tr.right.Value2())
	return tr
}

func (tr *Tree) GetValue1(l int, r int, L int, R int) int {
	if tr == nil {
		return inf
	}

	if L == l && r == R {
		return tr.val1
	}
	mid := (l + r) / 2
	ans := inf
	if L <= mid {
		ans = min(ans, tr.left.GetValue1(l, mid, L, min(mid, R)))
	}
	if mid < R {
		ans = min(ans, tr.right.GetValue1(mid+1, r, max(L, mid+1), r))
	}
	return ans
}
func (tr *Tree) GetValue2(l int, r int, L int, R int) int {
	if tr == nil {
		return -inf
	}
	if L == l && r == R {
		return tr.val2
	}
	mid := (l + r) / 2
	ans := -inf
	if L <= mid {
		ans = max(ans, tr.left.GetValue2(l, mid, L, min(mid, R)))
	}
	if mid < R {
		ans = max(ans, tr.right.GetValue2(mid+1, r, max(L, mid+1), R))
	}
	return ans
}
