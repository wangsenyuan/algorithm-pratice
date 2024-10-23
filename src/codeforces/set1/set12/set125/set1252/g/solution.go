package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) []int {
	n, m, k := readThreeNums(reader)

	A := readNNums(reader, n)
	B := make([][]int, m)
	for i := 0; i < m; i++ {
		var cnt int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &cnt) + 1
		B[i] = make([]int, cnt)
		for j := 0; j < cnt; j++ {
			pos = readInt(s, pos, &B[i][j]) + 1
		}
	}
	C := make([][]int, k)
	for i := 0; i < k; i++ {
		C[i] = readNNums(reader, 3)
	}

	return solve(A, B, C)

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

func solve(A []int, stages [][]int, queries [][]int) []int {
	m := len(stages)

	first := A[0]
	var over int
	for _, x := range A {
		if x >= first {
			over++
		}
	}
	n := len(A)
	rem := make([]int, m)
	for j := 0; j < m; j++ {
		rem[j] = n - over - len(stages[j])
		for _, x := range stages[j] {
			if x > first {
				over++
			}
		}
	}
	set := NewSegTree(rem)

	ok := set.Get(0, m-1) >= 0

	ans := make([]int, len(queries))

	for i, cur := range queries {
		x, y, z := cur[0], cur[1], cur[2]
		x--
		y--

		if stages[x][y] < first && first < z {
			set.Update(x+1, m-1, -1)
			ok = set.Get(0, m-1) >= 0
		} else if stages[x][y] > first && first > z {
			set.Update(x+1, m-1, 1)
			ok = set.Get(0, m-1) >= 0
		}

		stages[x][y] = z

		if ok {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}
	return ans
}

type SegTree struct {
	arr  []int
	lazy []int
	sz   int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		val[i] = min(val[2*i+1], val[2*i+2])
	}
	build(0, 0, n-1)
	return &SegTree{val, lazy, n}
}

func (tr *SegTree) update(i int, v int) {
	tr.arr[i] += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(i*2+1, tr.lazy[i])
		tr.update(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) pull(i int) {
	tr.arr[i] = min(tr.arr[i*2+1], tr.arr[i*2+2])
}

func (tr *SegTree) Update(L int, R int, v int) {
	if L > R {
		return
	}
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
		tr.pull(i)
	}
	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return 1 << 30
		}
		if L <= l && r <= R {
			return tr.arr[i]
		}
		mid := (l + r) / 2
		tr.push(i)
		a := loop(i*2+1, l, mid)
		b := loop(i*2+2, mid+1, r)
		return min(a, b)
	}
	return loop(0, 0, tr.sz-1)
}
