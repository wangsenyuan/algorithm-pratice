package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}
	buf.WriteTo(os.Stdout)
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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = readNNums(reader, 2)
	}
	return solve(a, ops)
}

func solve(a []int, ops [][]int) []int {
	// n := len(a)

	tr := NewTree(a)
	ans := make([]int, len(ops)+1)
	ans[0] = tr.val[0]

	for i, cur := range ops {
		tr.Update(cur[0]-1, cur[1])
		ans[i+1] = tr.val[0]
	}

	return ans
}

const inf = 1 << 50

type Tree struct {
	val0 []int // a[i] - i 的最小值
	val1 []int // a[i] - i 的最大值
	val2 []int // a[i] + i 的最小值
	val3 []int // a[i] + i 的最大值
	val  []int // 能够得到的最大值
	sz   int
}

func (tr *Tree) merge(i int) {
	tr.val0[i] = min(tr.val0[2*i+1], tr.val0[2*i+2])
	tr.val1[i] = max(tr.val1[2*i+1], tr.val1[2*i+2])
	tr.val2[i] = max(tr.val2[2*i+1], tr.val2[2*i+2])
	tr.val3[i] = min(tr.val3[2*i+1], tr.val3[2*i+2])

	tr.val[i] = max(tr.val[2*i+1], tr.val[2*i+2])
	// a[r] - r - (a[l] - l)
	tr.val[i] = max(tr.val[i], tr.val1[2*i+2]-tr.val0[2*i+1])
	// a[l] + l - (a[r] + r)
	tr.val[i] = max(tr.val[i], tr.val2[2*i+1]-tr.val3[2*i+2])
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	tr := new(Tree)
	tr.val0 = make([]int, 4*n)
	tr.val1 = make([]int, 4*n)
	tr.val2 = make([]int, 4*n)
	tr.val3 = make([]int, 4*n)
	tr.val = make([]int, 4*n)

	tr.sz = n

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			tr.val0[i] = arr[l] - l
			tr.val1[i] = arr[l] - l
			tr.val2[i] = arr[l] + l
			tr.val3[i] = arr[l] + l
			tr.val[i] = 0
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		tr.merge(i)
	}
	build(0, 0, n-1)
	return tr
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val0[i] = v - l
			tr.val1[i] = v - l
			tr.val2[i] = v + l
			tr.val3[i] = v + l
			tr.val[i] = 0
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.merge(i)
	}
	loop(0, 0, tr.sz-1)
}
