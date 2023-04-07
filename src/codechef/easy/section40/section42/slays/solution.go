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

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		S := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(S, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const INF = 1e9

const N_INF = -INF

func solve(S []int, Q [][]int) []int {
	// 如果 D >= max(S[l...r]) + 1
	// 显然可以完成目标
	// D = max(S[l...r] 是否ok？ (1)
	// 首先不能有2个max
	// 然后D, D-1, D - 2, ... D-i, 假设有这样一个连续的下降的路径 (2)
	// 后面i+1..r 必须 < D - i - 1 (3)
	// 1 和 3 可以使用segment tree
	// 但是2的话，以此迭代肯定是不行的，但是如果，
	// 把下标倒过来，从后往前，D, D - 1, D - 2, D - i... 连续的等于(D - x) （假设D的下标是x）
	// 且 Max（i + 1...r) < D - x
	// 似乎也不对
	//  D ... D-1, ... D-2, ... D-x
	//  D... D-1 中间的数 < D - 1, 就不影响结果
	// 那么D就是ok的
	// 咋处理呢？
	n := len(S)
	root := NewTree(n, S)

	pos := make(map[int]int)

	for i := 0; i < n; i++ {
		pos[S[i]] = i
	}

	arr := make([]int, 0, len(pos))

	for k := range pos {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	// R[i] = right-position of i, having S[j] - (n - 1 - j) == S[i] - (n - 1 - i)
	R := make([]int, n)
	R[n-1] = n - 1
	s := NewSet(n)

	for i := n - 1; i >= 0; i-- {
		d := S[i]
		j := search(len(arr), func(j int) bool {
			return arr[j] >= d-1
		})
		k := s.Get(j, len(arr))
		if k >= n || S[k] >= d {
			R[i] = min(n, k) - 1
		} else {
			R[i] = R[k]
		}
		j = search(len(arr), func(j int) bool {
			return arr[j] >= d
		})
		s.Set(j, i)
	}

	res := make([]int, len(Q))

	for i, cur := range Q {
		l, r := cur[0], cur[1]
		l--
		r--
		// 找到这个区间内的最大值
		node := root.Get(l, r)
		res[i] = node.val + 1
		if node.cnt > 1 {
			continue
		}
		p := node.pos
		if R[p] >= r {
			res[i] = node.val
		}
	}

	return res
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

type Seg struct {
	arr []int
	sz  int
}

func NewSet(n int) *Seg {
	arr := make([]int, 2*n)

	for i := 0; i < 2*n; i++ {
		arr[i] = INF
	}

	return &Seg{arr, n}
}

func (s *Seg) Set(p int, v int) {
	p += s.sz
	s.arr[p] = v
	for p > 1 {
		s.arr[p>>1] = min(s.arr[p], s.arr[p^1])
		p >>= 1
	}
}

func (s *Seg) Get(l int, r int) int {
	l += s.sz
	r += s.sz

	var res int = INF

	for l < r {
		if l&1 == 1 {
			res = min(res, s.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, s.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type Item struct {
	val int
	cnt int
	pos int
}

type Tree struct {
	items []*Item
	sz    int
}

func NewTree(n int, arr []int) *Tree {
	items := make([]*Item, 4*n)

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		items[i] = new(Item)
		if l == r {
			items[i].pos = l
			items[i].val = arr[l]
			items[i].cnt = 1
			return
		}
		mid := (l + r) / 2
		build(2*i, l, mid)
		build(2*i+1, mid+1, r)
		pull(items[i], items[2*i], items[2*i+1])
	}

	build(1, 0, n-1)

	sz := n
	return &Tree{items, sz}
}

func pull(a, b, c *Item) {
	if b.val > c.val {
		a.val = b.val
		a.cnt = b.cnt
		a.pos = b.pos
		return
	}
	a.val = c.val
	a.cnt = c.cnt
	a.pos = c.pos
	if b.val == c.val {
		a.cnt += b.cnt
	}
}

func (t *Tree) Get(L int, R int) *Item {
	var loop func(i int, l int, r int) *Item
	loop = func(i int, l int, r int) *Item {
		res := new(Item)
		res.val = N_INF
		if R < l || r < L {
			return res
		}
		if L <= l && r <= R {
			return t.items[i]
		}
		mid := (l + r) / 2
		a := loop(2*i, l, mid)
		b := loop(2*i+1, mid+1, r)
		pull(res, a, b)
		return res
	}
	return loop(1, 0, t.sz-1)
}
