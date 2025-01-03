package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, r := readTwoNums(reader)
	points := make([][]int, n)
	for i := range n {
		points[i] = readNNums(reader, 2)
	}
	return solve(points, r)
}

const inf = 2 * 1e6

func solve(points [][]int, r int) int {
	lines := make(map[int][]int)
	bounds := []int{inf, -inf}
	for _, cur := range points {
		x, y := cur[0], cur[1]
		if lines[x+y] == nil {
			lines[x+y] = make([]int, 0, 2)
		}
		lines[x+y] = append(lines[x+y], y-x)
		bounds[0] = min(bounds[0], y-x)
		bounds[1] = max(bounds[1], y+x)
	}

	var arr []int
	for k := range lines {
		arr = append(arr, k)
		sort.Ints(lines[k])
	}

	sort.Ints(arr)
	var res int

	m := bounds[1] - bounds[0] + 1
	tr := NewTree(m)

	r *= 2

	for i, j := 0, 0; i < len(arr); i++ {
		x := arr[i]
		for j < len(arr) && arr[j] <= x+r {
			for _, c := range lines[arr[j]] {
				tr.Update(max(0, c-r-bounds[0]), c-bounds[0], 1)
				cnt := tr.Get(0, m-1)
				res = max(res, cnt)
			}
			j++
		}
		for _, c := range lines[x] {
			tr.Update(max(0, c-r-bounds[0]), c-bounds[0], -1)
		}
	}

	return res
}

type Tree struct {
	arr  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{arr, lazy, n}
}

func (tr *Tree) update(i int, v int) {
	tr.arr[i] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		tr.arr[i] = max(tr.arr[2*i+1], tr.arr[2*i+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return tr.arr[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		return max(loop(2*i+1, l, mid), loop(2*i+2, mid+1, r))
	}
	return loop(0, 0, tr.sz-1)
}
