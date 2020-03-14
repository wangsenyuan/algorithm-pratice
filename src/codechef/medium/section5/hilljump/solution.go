package main

import (
	"bufio"
	"fmt"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64s(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	pos := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for pos < len(scanner.Bytes()) && scanner.Bytes()[pos] == ' ' {
			pos++
		}
		pos = readInt64(scanner.Bytes(), pos, &res[i])
	}
	return res
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	A := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &A[i])
	}

	solver := NewSolver1(n, A)

	// for i := (n - 1) / BL; i >= 0; i-- {
	// 	solver.UpdateBlock(i)
	// }

	for m > 0 {
		m--
		var tc int
		fmt.Scanf("%d", &tc)
		if tc == 1 {
			var i, k int
			fmt.Scanf("%d %d", &i, &k)
			i--
			fmt.Println(solver.Query(i, k) + 1)
		} else {
			var l, r, x int
			fmt.Scanf("%d %d %d", &l, &r, &x)
			l--
			r--
			solver.EditRange(l, r, int64(x))
		}

	}
}

type Solver1 struct {
	n     int
	A     []int64
	B     *BIT
	jumps []Jump
	buf   *Buffer
}

func NewSolver1(n int, A []int64) Solver1 {
	B := NewBIT(n)
	jumps := make([]Jump, n)

	buf := NewBuffer(BL)

	update(n, A, &B, jumps, &buf, 0, n)

	return Solver1{n, A, &B, jumps, &buf}
}

func update(n int, A []int64, B *BIT, jumps []Jump, buf *Buffer, l, r int) {
	buf.Reset()
	for i := min(n, r+BL) - 1; i >= r; i-- {
		ai := A[i] + B.Get(i)
		for !buf.IsEmpty() && ai >= buf.Back().val {
			buf.PopBack()
		}
		buf.PushBack(Item{i, ai, jumps[i]})
	}

	for i := r - 1; i >= l; i-- {
		ai := A[i] + B.Get(i)
		for !buf.IsEmpty() && ai >= buf.Back().val {
			buf.PopBack()
		}
		if buf.IsEmpty() {
			jumps[i] = Jump{i, 0}
		} else {
			if buf.Back().jump.pos-i > BL {
				// far away
				jumps[i] = Jump{buf.Back().pos, 1}
			} else {
				jumps[i] = buf.Back().jump.Clone()
				jumps[i].IncrSteps()
			}

			if buf.Front().pos-i >= BL {
				buf.PopFront()
			}
		}
		buf.PushBack(Item{i, ai, jumps[i]})
	}
}

func (solver Solver1) EditRange(l, r int, x int64) {
	r++
	solver.B.Update(l, x)
	solver.B.Update(r, -x)
	if r-l < 2*BL {
		update(solver.n, solver.A, solver.B, solver.jumps, solver.buf, max(l-BL, 0), r)
	} else {
		update(solver.n, solver.A, solver.B, solver.jumps, solver.buf, max(l-BL, 0), l)
		update(solver.n, solver.A, solver.B, solver.jumps, solver.buf, r-BL, r)
	}
}

func (solver Solver1) Query(i, k int) int {
	for solver.jumps[i].steps > 0 && solver.jumps[i].steps <= k {
		k -= solver.jumps[i].steps
		i = solver.jumps[i].pos
	}

	if k > 0 && solver.jumps[i].steps > 0 {
		ai := solver.A[i] + solver.B.Get(i)
		for j := i + 1; j < min(i+BL+1, solver.n)-1; j++ {
			aj := solver.A[j] + solver.B.Get(j)
			if aj > ai {
				i = j
				ai = aj
				k--
				if k == 0 {
					break
				}
			}
		}
	}
	return i
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type BIT struct {
	arr []int64
	n   int
}

func NewBIT(n int) BIT {
	arr := make([]int64, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Update(pos int, v int64) {
	pos++

	for pos <= bit.n {
		bit.arr[pos] += v
		pos += pos & -pos
	}
}

func (bit BIT) Get(pos int) int64 {
	pos++
	var res int64
	for pos > 0 {
		res += bit.arr[pos]
		pos -= pos & -pos
	}
	return res
}

type Jump struct {
	pos   int
	steps int
}

func (this Jump) Clone() Jump {
	return Jump{this.pos, this.steps}
}

func (this *Jump) IncrSteps() {
	this.steps++
}

type Item struct {
	pos  int
	val  int64
	jump Jump
}

type Buffer struct {
	items []Item
	size  int
	head  int
	tail  int
}

func NewBuffer(size int) Buffer {
	// avoid full/empty check conflicts
	size++
	items := make([]Item, size)
	return Buffer{items, size, 0, 0}
}

func (buffer Buffer) IsEmpty() bool {
	return buffer.head == buffer.tail
}

func (buffer *Buffer) PushBack(item Item) {
	buffer.items[buffer.tail] = item
	buffer.tail = (buffer.tail + 1) % buffer.size
}

func (buffer Buffer) Front() Item {
	return buffer.items[buffer.head]
}

func (buffer Buffer) Back() Item {
	tail := buffer.tail - 1
	if tail < 0 {
		tail += buffer.size
	}
	return buffer.items[tail]
}

func (buffer *Buffer) PopFront() Item {
	item := buffer.items[buffer.head]
	buffer.head = (buffer.head + 1) % buffer.size
	return item
}

func (buffer *Buffer) PopBack() Item {
	tail := buffer.tail - 1
	if tail < 0 {
		tail += buffer.size
	}
	item := buffer.items[tail]
	buffer.tail = tail
	return item
}

func (buffer *Buffer) Reset() {
	buffer.head = 0
	buffer.tail = 0
}

type Solver struct {
	n     int
	A     []int64
	nxt   []int
	mask  []int64
	stops []int
	steps []int
	ind   []int
	st    []int64
	cnt   []int
}

const BL = 100

func NewSolver(n int, A []int64) Solver {
	nxt := make([]int, n)
	mask := make([]int64, n/BL+3)
	stops := make([]int, n)
	steps := make([]int, n)
	ind := make([]int, n)
	st := make([]int64, n)
	cnt := make([]int, n)
	return Solver{n, A, nxt, mask, stops, steps, ind, st, cnt}
}

func (solver *Solver) UpdateBlock(x int) {
	if x < 0 {
		return
	}
	if solver.mask[x] != 0 {
		for i := BL * x; i < min(solver.n, BL*(x+1)); i++ {
			solver.A[i] += solver.mask[x]
		}
		solver.mask[x] = 0
	}
	if solver.mask[x+1] != 0 {
		for i := BL * (x + 1); i < min(solver.n, BL*(x+2)); i++ {
			solver.A[i] += solver.mask[x+1]
		}
		solver.mask[x+1] = 0
	}

	var si int
	for i := min(solver.n, (x+2)*BL) - 1; i >= (x+1)*BL; i-- {
		for si > 0 && solver.A[i] >= solver.st[si-1] {
			si--
		}
		solver.st[si] = solver.A[i]
		solver.steps[si] = 0
		solver.stops[si] = i
		solver.ind[si] = i
		si++
	}

	for i := min(solver.n, (x+1)*BL) - 1; i >= x*BL; i-- {
		for si > 0 && solver.A[i] >= solver.st[si-1] {
			si--
		}
		if si == 0 || solver.ind[si-1]-i > BL {
			solver.cnt[i] = 0
			solver.nxt[i] = -1
		} else {
			solver.st[si] = solver.A[i]
			solver.steps[si] = solver.steps[si-1] + 1
			solver.stops[si] = solver.stops[si-1]
			solver.ind[si] = i

			solver.cnt[i] = solver.steps[si]
			solver.nxt[i] = solver.stops[si]
			si++
		}
	}
}

func (solver *Solver) Query(place int, remain int) int {
	for {
		if solver.cnt[place] > remain {
			break
		}
		if solver.cnt[place] == 0 {
			break
		}
		remain -= solver.cnt[place]
		place = solver.nxt[place]
	}

	if remain == 0 {
		return place
	}

	cur := place / BL

	for i := place + 1; i < min(solver.n, (cur+1)*BL); i++ {
		if remain == 0 {
			break
		}
		if solver.A[i] > solver.A[place] {
			place = i
			remain--
		}
	}
	return place
}

func (solver *Solver) EditRange(l, r int, x int64) {
	if l/BL == r/BL {
		for i := l; i <= r; i++ {
			solver.A[i] += x
		}
		xx := l / BL
		solver.UpdateBlock(xx)
		solver.UpdateBlock(xx - 1)
		return
	}
	rbl := r / BL
	lbl := l / BL
	for i := rbl * BL; i <= r; i++ {
		solver.A[i] += x
	}

	for i := l; i < (lbl+1)*BL; i++ {
		solver.A[i] += x
	}

	for i := lbl + 1; i < rbl; i++ {
		solver.mask[i] += x
	}

	solver.UpdateBlock(rbl)
	solver.UpdateBlock(rbl - 1)

	if lbl < rbl-1 {
		solver.UpdateBlock(lbl)
	}
	solver.UpdateBlock(lbl - 1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
