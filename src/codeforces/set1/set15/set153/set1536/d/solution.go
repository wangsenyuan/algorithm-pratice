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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(b []int) bool {
	b = compact(b)
	n := len(b)
	tr := NewSegTree(n)
	for i := 0; i < n; i++ {
		tr.Update(i, -1)
	}

	tr.Update(b[0], b[0])

	for i := 1; i < n; i++ {
		if b[i-1] < b[i] {
			x := tr.Query(b[i-1]+1, b[i])
			if x >= 0 {
				return false
			}
		} else if b[i-1] > b[i] {
			x := tr.Query(b[i]+1, b[i-1])
			if x >= 0 {
				return false
			}
		}
		tr.Update(b[i], b[i])
	}

	return true
}

func compact(arr []int) []int {
	id := make(map[int]int)
	for _, num := range arr {
		id[num]++
	}
	var nums []int
	for k := range id {
		nums = append(nums, k)
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		id[nums[i]] = i
	}
	res := make([]int, len(arr))
	for i, num := range arr {
		res[i] = id[num]
	}
	return res
}

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	return SegTree(arr)
}

func (tree SegTree) Update(p int, v int) {
	n := len(tree)
	n /= 2
	p += n
	tree[p] = v
	for p > 0 {
		tree[p>>1] = max(tree[p], tree[p^1])
		p >>= 1
	}
}

func (tree SegTree) Query(l, r int) int {
	n := len(tree)
	n /= 2
	l += n
	r += n
	res := -1
	for l < r {
		if l&1 == 1 {
			res = max(res, tree[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tree[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
