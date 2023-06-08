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
		n := readNum(reader)
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(reader, 2)
		}
		res := solve(segs)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(segs [][]int) int {
	n := len(segs)
	items := make([]Item, n)
	for i := 0; i < n; i++ {
		items[i] = Item{segs[i][0], segs[i][1], i}
	}

	// R < l and r < L
	ans := make([]int, n)
	sort.Slice(items, func(i, j int) bool {
		return items[i].r < items[j].r
	})

	for _, it := range items {
		j := search(n, func(j int) bool {
			return items[j].r >= it.l
		})
		j--
		// items[j].r < it.l
		ans[it.id] = j + 1
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].l < items[j].l
	})

	for i := n - 1; i >= 0; i-- {
		it := items[i]
		j := search(n, func(j int) bool {
			return items[j].l > it.r
		})
		ans[it.id] += n - j
	}

	best := n
	for i := 0; i < n; i++ {
		best = min(best, ans[i])
	}
	return best
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

func solve1(segs [][]int) int {
	// 如果删除到只剩一个，肯定是可以的
	compress := make(map[int]int)
	for _, seg := range segs {
		compress[seg[0]]++
		compress[seg[1]]++
	}
	arr := make([]int, 0, len(compress))
	for k := range compress {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		compress[arr[i]] = i
	}

	n := len(segs)
	items := make([]Item, n)
	for i := 0; i < n; i++ {
		items[i] = Item{compress[segs[i][0]], compress[segs[i][1]], i}
	}
	// 陷入了漩涡，休息一下。
	// 假设选中seg x做为特殊的, 那么就需要找到和它能相交的部分
	// 包含4部分， 1. 完全包含在它里面的部分
	//           2. 只和它的左端点相交的部分
	//           3. 只和它的右端点相交的部分
	//           4. 完全包含它的部分
	// 1和4是相对的， 有没有办法可以很快的计算出来 （第4个其实不用算，因为完全包含它的会有个至少不必它差的结果）

	ans := make([]int, n)

	// 计算 2 & 4，右端点处在l..r中间的部分
	sort.Slice(items, func(i, j int) bool {
		return items[i].r < items[j].r || items[i].r == items[j].r && items[i].l > items[j].l
	})

	m := len(arr)
	active := NewSegTree(m)

	for _, it := range items {
		i, l, r := it.id, it.l, it.r
		ans[i] = active.Get(l, r+1)
		active.Update(r, 1)
	}

	active.Reset()
	// 然后找出右端点在r后面，但是左端点在r(包括)前面的
	sort.Slice(items, func(i, j int) bool {
		return items[i].r > items[j].r || items[i].r == items[j].r && items[i].l < items[j].l
	})

	for i := 0; i < n; {
		tmp := active.Get(0, items[i].r+1)
		j := i
		for i < n && items[i].r == items[j].r {
			ans[items[i].id] += tmp
			active.Update(items[i].l, 1)
			i++
		}
	}
	best := n
	for i := 0; i < n; i++ {
		best = min(best, n-ans[i]-1)
	}

	return best
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] += v
	for p > 1 {
		t.arr[p>>1] = t.arr[p] + t.arr[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	l += t.n
	r += t.n
	var res int
	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func (t *SegTree) Reset() {
	for i := 0; i < len(t.arr); i++ {
		t.arr[i] = 0
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	l  int
	r  int
	id int
}

type Event struct {
	id  int
	pos int
	tp  int
}
