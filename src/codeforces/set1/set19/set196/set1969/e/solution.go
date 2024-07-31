package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
func solve(a []int) int {
	n := len(a)
	tr := Create(n)

	var st int
	var res int

	pos := make([][]int, n)

	for i := 0; i < n; i++ {
		x := a[i] - 1
		pos[x] = append(pos[x], i)
		k := len(pos[x])
		// 这个顺序很重要
		// 在加入x前，st后面所有的数字，应该是只出现了一次
		// 所以，k > 0 这里，只会增加上去
		// 如果x出现了两次，这个时候倒数第二个x的值（包括它前面的）至少是2
		// 所以这里-2， 会让其变成0 （但如果这两个位置中间正好出现了一个新的）似乎也不会呐
		if k > 0 {
			tr.Update(st, pos[x][k-1], 1)
		}
		if k > 1 {
			tr.Update(st, pos[x][k-2], -2)
		}
		if k > 2 {
			// 再前面的要变成0
			tr.Update(st, pos[x][k-3], 1)
		}
		if tr.Get(st, i) == 0 {
			res++
			st = i + 1
		}
	}

	return res
}

type SegTree struct {
	lazy []int
	val  []int
	sz   int
}

func Create(n int) *SegTree {
	lazy := make([]int, 4*n)
	val := make([]int, 4*n)
	return &SegTree{lazy, val, n}
}

func (tr *SegTree) pushValue(i int, v int) {
	tr.lazy[i] += v
	tr.val[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.pushValue(i*2+1, tr.lazy[i])
		tr.pushValue(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			tr.pushValue(i, v)
			return
		}
		tr.push(i)

		mid := (l + r) / 2
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		tr.val[i] = min(tr.val[2*i+1], tr.val[i*2+2])
	}

	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int

	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return 1 << 50
		}
		if L <= l && r <= R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		a := loop(i*2+1, l, mid)
		b := loop(i*2+2, mid+1, r)
		return min(a, b)
	}

	return loop(0, 0, tr.sz-1)
}
