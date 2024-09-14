package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	lectures := make([][]int, n)
	for i := 0; i < n; i++ {
		lectures[i] = readNNums(reader, 4)
	}
	res := solve(lectures)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

type pair struct {
	first  int
	second int
}

func construct(lectures [][]int, id int) []pair {
	n := len(lectures)
	res := make([]pair, n)
	for i := 0; i < n; i++ {
		res[i] = pair{lectures[i][id*2], lectures[i][id*2+1]}
	}

	return res
}

func getDistinct(arr []pair) []int {
	var nums []int
	for _, cur := range arr {
		nums = append(nums, cur.first, cur.second)
	}
	sort.Ints(nums)

	var res []int

	for i := 0; i < len(nums); {
		res = append(res, nums[i])
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
	}

	return res
}

func solve(lectures [][]int) bool {
	if process(lectures, 0) || process(lectures, 1) {
		// 存在 sensitive set
		return false
	}
	return true
}

func process(lectures [][]int, first int) bool {
	a := construct(lectures, first)
	b := construct(lectures, first^1)
	pos := getDistinct(b)

	n := len(b)
	for i := 0; i < n; i++ {
		x := sort.SearchInts(pos, b[i].first)
		y := sort.SearchInts(pos, b[i].second)
		// rewrite by position
		b[i] = pair{x, y}
	}

	events := make([]pair, 2*n)

	for i, cur := range a {
		events[2*i] = pair{cur.first, i + 1}
		events[2*i+1] = pair{cur.second + 1, -(i + 1)}
	}

	slices.SortFunc(events, func(x, y pair) int {
		if x.first != y.first {
			return x.first - y.first
		}
		// 先处理结束事件
		return x.second - y.second
	})

	set := NewSegTree(len(pos))

	var active int

	for _, cur := range events {
		if cur.second < 0 {
			// leave
			i := -(cur.second) - 1
			x, y := b[i].first, b[i].second
			set.Update(x, y, -1)
			active--
			continue
		}

		i := cur.second - 1
		x, y := b[i].first, b[i].second
		// 看这个区间内，是否没有重叠的
		v := set.Get(x, y)
		if v != active {
			// 肯定有一个区间和它没有重叠
			return true
		}
		set.Update(x, y, 1)
		active++
	}

	return false
}

type SegTree struct {
	arr  []int
	lazy []int
	n    int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &SegTree{arr, lazy, n}
}

func (tr *SegTree) update(i int, v int) {
	tr.arr[i] += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) pull(i int) {
	tr.arr[i] = max(tr.arr[i*2+1], tr.arr[i*2+2])
}
func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if r < L || R < l {
			return
		}
		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		tr.pull(i)
	}

	loop(0, 0, tr.n-1)
}

func (tr *SegTree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return 0
		}
		if L <= l && r <= R {
			return tr.arr[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		x := loop(i*2+1, l, mid)
		y := loop(i*2+2, mid+1, r)
		return max(x, y)
	}

	return loop(0, 0, tr.n-1)
}
