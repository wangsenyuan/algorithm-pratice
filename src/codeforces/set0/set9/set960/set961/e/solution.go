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
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

type BIT []int

func (bit BIT) Add(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) Query(p int) int {
	p++
	res := 0
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func solve(a []int) int {
	n := len(a)
	bit := make(BIT, n+10)
	arr := make([][]int, n)
	var res int

	for i := range n {
		a[i]--
		x := a[i]
		if x < n {
			arr[x] = append(arr[x], i)
		}
		bit.Add(i, 1)
		if i <= x {
			res--
		}
	}

	for i := 0; i < n; i++ {
		j := min(n-1, a[i])
		res += bit.Query(j)

		for _, k := range arr[i] {
			bit.Add(k, -1)
		}
	}

	return res / 2
}

type pair struct {
	first, second int
}

func solve1(a []int) int {
	n := len(a)

	// 对于(i, j)来说 => i <= a[j] & j <= a[i]的计数
	// 对于pair (a[i], i) 来说，就是找到这样的j (a[j] >= i, j <= a[i])
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{min(a[i], n), i + 1}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return b.first - a.first
	})

	trs := make([]*Tree, n)
	for i, cur := range arr {
		if i == 0 {
			trs[i] = trs[i].Add(1, n, cur.second)
		} else {
			trs[i] = trs[i-1].Add(1, n, cur.second)
		}
	}

	var res int

	for _, cur := range arr {
		x := cur.first
		i := cur.second
		j := sort.Search(n, func(j int) bool {
			return arr[j].first < i
		})
		j--
		if j < 0 {
			// 没有比i大的
			continue
		}
		// arr[j].first >= i
		cnt := trs[j].Query(1, n, x)
		res += cnt
		if i <= x {
			// 要把自己的贡献去掉
			res--
		}
	}

	return res / 2
}

type Tree struct {
	left, right *Tree
	cnt         int
}

func (tr *Tree) Count() int {
	if tr == nil {
		return 0
	}
	return tr.cnt
}

func (tr *Tree) Add(l, r int, p int) *Tree {
	res := new(Tree)

	if tr != nil {
		res.left = tr.left
		res.right = tr.right
		res.cnt = tr.cnt
	}

	res.cnt++

	if l < r {
		mid := (l + r) / 2
		if p <= mid {
			res.left = res.left.Add(l, mid, p)
		} else {
			res.right = res.right.Add(mid+1, r, p)
		}
	}

	return res
}

func (tr *Tree) Query(l int, r int, x int) int {
	if tr == nil {
		return 0
	}
	if l == r {
		return tr.cnt
	}
	mid := (l + r) / 2
	if x <= mid {
		return tr.left.Query(l, mid, x)
	}
	return tr.left.Count() + tr.right.Query(mid+1, r, x)
}
