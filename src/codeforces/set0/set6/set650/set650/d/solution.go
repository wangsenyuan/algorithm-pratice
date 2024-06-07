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

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	res := solve(a, queries)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

type Pair struct {
	first  int
	second int
}

func solve(h []int, queries [][]int) []int {
	n := len(h)

	lis, p1 := findPosition(h)
	var p2 []int
	{
		_h := make([]int, n)
		copy(_h, h)
		for i := 0; i < n; i++ {
			_h[i] *= -1
		}
		reverse(_h)
		_, p2 = findPosition(_h)
		reverse(p2)
	}
	uni := make([]bool, n)
	cnt := make([]int, n+1)

	for i := 0; i < n; i++ {
		if p1[i]+p2[i]-1 == lis {
			cnt[p1[i]]++
		}
	}

	for i := 0; i < n; i++ {
		uni[i] = cnt[p1[i]] == 1 && p1[i]+p2[i]-1 == lis
	}

	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{h[i], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first || arr[i].first == arr[j].first && arr[i].second < arr[j].second
	})

	getPos := func(p Pair) int {
		return search(n, func(i int) bool {
			return arr[i].first > p.first || arr[i].first == p.first && arr[i].second >= p.second
		})
	}

	next := NewTree(n)

	for i := 0; i < n; i++ {
		p := Pair{h[i], i}
		j := getPos(p)
		next.Update(j, p2[i])
	}

	prev := NewTree(n)

	// 将a[i]变大以后，理论上它会增加左边的贡献
	// 所以，应该找到左边 < x中p1[j]最大的j，
	// 右边 > x中p2[k]最大的k

	apply := func(i int, x int) int {
		if x == h[i] {
			return lis
		}

		j := getPos(Pair{x, -1})
		l := prev.Get(0, j)
		j = getPos(Pair{x, n})
		r := next.Get(j, n)

		if uni[i] {
			return max(l+1+r, lis-1)
		}

		return max(lis, l+r+1)
	}

	ans := make([]int, len(queries))

	at := make([][]int, n)
	for i, cur := range queries {
		j := cur[0] - 1
		at[j] = append(at[j], i)
	}

	for j := 0; j < n; j++ {
		it := getPos(Pair{h[j], j})
		next.Update(it, 0)

		for _, i := range at[j] {
			cur := queries[i]
			ans[i] = apply(j, cur[1])
		}

		prev.Update(it, p1[j])
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func findPosition(a []int) (r int, pos []int) {
	n := len(a)
	nums := make([]int, n)
	copy(nums, a)
	// pos[i] = i 在LIS中的前序位置
	pos = make([]int, n)

	for i := 0; i < n; i++ {
		j := search(r, func(j int) bool {
			return nums[j] >= nums[i]
		})
		pos[i] = j + 1
		nums[j] = nums[i]
		if j == r {
			r++
		}
	}
	return r, pos
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

type Tree struct {
	arr []int
	sz  int
}

func NewTree(n int) *Tree {
	arr := make([]int, 2*n)
	return &Tree{arr, n}
}

func (tr *Tree) Update(p int, v int) {
	p += tr.sz
	tr.arr[p] = v

	for p > 1 {
		tr.arr[p>>1] = max(tr.arr[p], tr.arr[p^1])
		p >>= 1
	}
}

func (tr *Tree) Get(l int, r int) int {
	l += tr.sz
	r += tr.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
