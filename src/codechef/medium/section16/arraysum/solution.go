package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, m)

	res := solve(A, B)

	fmt.Println(res)
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

const MAX_X = 1000000000

func solve(A []int, B []int) int {
	n := len(A)
	m := len(B)

	const BLK_SIZE = 10000

	var left [][]int
	var right [][]int
	var val [][]int

	var ptr int
	var sz int
	next := func() int {
		if ptr == sz {
			left = append(left, make([]int, BLK_SIZE))
			right = append(right, make([]int, BLK_SIZE))
			val = append(val, make([]int, BLK_SIZE))
			j := len(left)
			for i := 0; i < BLK_SIZE; i++ {
				left[j-1][i] = -1
				right[j-1][i] = -1
			}

			sz += BLK_SIZE
		}
		ptr++
		return ptr - 1
	}

	getVal := func(x int) int {
		i, j := x/BLK_SIZE, x%BLK_SIZE
		return val[i][j]
	}

	getLeft := func(x int) int {
		i, j := x/BLK_SIZE, x%BLK_SIZE
		return left[i][j]
	}

	getRight := func(x int) int {
		i, j := x/BLK_SIZE, x%BLK_SIZE
		return right[i][j]
	}

	setLeft := func(x int, v int) {
		i, j := x/BLK_SIZE, x%BLK_SIZE
		left[i][j] = v
	}

	setRight := func(x int, v int) {
		i, j := x/BLK_SIZE, x%BLK_SIZE
		right[i][j] = v
	}

	setVal := func(x int, v int) {
		i, j := x/BLK_SIZE, x%BLK_SIZE
		val[i][j] = v
	}

	cloneNode := func(i int) int {
		j := next()
		setLeft(j, getLeft(i))
		setRight(j, getRight(i))
		setVal(j, getVal(i))
		return j
	}

	push := func(x int) {
		if getLeft(x) < 0 {
			setLeft(x, next())
			setRight(x, next())
		}
	}

	merge := func(x int) {
		setVal(x, modAdd(getVal(getLeft(x)), getVal(getRight(x))))
	}

	var add func(i int, begin int, end int, p int, v int) int

	add = func(i int, begin int, end int, p int, v int) int {
		j := cloneNode(i)

		if begin == end {
			setVal(j, modAdd(getVal(j), v))
		} else {
			push(j)
			mid := (begin + end) / 2
			if p <= mid {
				setLeft(j, add(getLeft(j), begin, mid, p, v))
			} else {
				setRight(j, add(getRight(j), mid+1, end, p, v))
			}
			merge(j)
		}
		return j
	}

	var query func(i int, begin int, end int, l int, r int) int

	query = func(i int, begin int, end int, l int, r int) int {
		if end < l || r < begin {
			return 0
		}
		if l <= begin && end <= r {
			return getVal(i)
		}
		var res int
		mid := (begin + end) / 2
		if getLeft(i) >= 0 {
			res = query(getLeft(i), begin, mid, l, r)
		}
		if getRight(i) >= 0 {
			res = modAdd(res, query(getRight(i), mid+1, end, l, r))
		}

		return res
	}

	trees := make([]int, n+1)
	trees[0] = next()
	// trees[0].Add(0, 0)

	for i := 1; i <= n; i++ {
		trees[i] = add(trees[i-1], 0, MAX_X, A[i-1], 1)
	}

	dp := make([]int, n+1)
	for j := 1; j < m; j++ {
		for i := j; i <= n; i++ {
			// A[i-1] + B[j] >= x + B[j-1]
			// x <= A[i-1] + B[j] - B[j-1]
			x := A[i-1] + B[j] - B[j-1]
			dp[i] = query(trees[i-1], 0, MAX_X, 0, x)
		}

		for i := j; i <= n; i++ {
			trees[i] = add(trees[i-1], 0, MAX_X, A[i-1], dp[i])
		}
	}

	var res int

	for i := m; i <= n; i++ {
		res = modAdd(res, dp[i])
	}

	return res
}

func solve1(A []int, B []int) int {
	n := len(A)
	m := len(B)

	trees := make([]*Node, n+1)
	trees[0] = NewNode(0, MAX_X)
	// trees[0].Add(0, 0)

	for i := 1; i <= n; i++ {
		trees[i] = trees[i-1].Add(A[i-1], 1)
	}

	dp := make([]int, n+1)
	for j := 1; j < m; j++ {
		for i := j; i <= n; i++ {
			// A[i-1] + B[j] >= x + B[j-1]
			// x <= A[i-1] + B[j] - B[j-1]
			x := A[i-1] + B[j] - B[j-1]
			dp[i] = trees[i-1].Query(0, x)
		}

		for i := j; i <= n; i++ {
			trees[i] = trees[i-1].Add(A[i-1], dp[i])
		}
	}

	var res int

	for i := m; i <= n; i++ {
		res = modAdd(res, dp[i])
	}

	return res
}

type Node struct {
	left, right *Node
	begin, end  int
	val         int
}

func NewNode(begin, end int) *Node {
	node := new(Node)
	node.begin = begin
	node.end = end
	return node
}

func (this *Node) push() {
	if this.left == nil {
		mid := (this.begin + this.end) / 2
		this.left = NewNode(this.begin, mid)
		this.right = NewNode(mid+1, this.end)
	}
}

func (this *Node) merge() {
	this.val = modAdd(this.left.val, this.right.val)
}

func (this *Node) copy() *Node {
	that := NewNode(this.begin, this.end)
	that.val = this.val
	that.left = this.left
	that.right = this.right
	return that
}

func (this *Node) Add(p int, v int) *Node {
	that := this.copy()

	if that.begin == that.end {
		that.val = modAdd(that.val, v)
	} else {
		that.push()
		mid := (that.begin + that.end) / 2
		if p <= mid {
			that.left = that.left.Add(p, v)
		} else {
			that.right = that.right.Add(p, v)
		}
		that.merge()
	}
	return that
}

func (this *Node) Query(l int, r int) int {
	if this == nil || this.end < l || r < this.begin {
		return 0
	}

	if l <= this.begin && this.end <= r {
		return this.val
	}

	res := this.left.Query(l, r)
	res = modAdd(res, this.right.Query(l, r))

	return res
}
