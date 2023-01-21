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
		n, q := readTwoNums(reader)
		A := readNNums(reader, n)
		solver := NewSolver(n, A)

		buf.WriteString(fmt.Sprintf("%d\n", solver.ans))

		for q > 0 {
			q--
			l, r := readTwoNums(reader)
			buf.WriteString(fmt.Sprintf("%d\n", solver.Swap(l, r)))
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

type Solver struct {
	n   int
	A   []int
	ans int64
}

func NewSolver(n int, A []int) Solver {
	x := make([]int, n+2)
	copy(x[1:], A)
	x[0] = -1
	x[n+1] = -1
	var ans int64

	for i := 1; i <= n; i++ {
		if x[i] > x[i-1] && x[i] > x[i+1] {
			ans += int64(x[i])
		} else if x[i] < x[i-1] && x[i] < x[i+1] {
			ans -= int64(x[i])
		}
	}

	return Solver{n, x, ans}
}

func (solver *Solver) Swap(l, r int) int64 {
	if l == r {
		return solver.ans
	}
	a := solver.A

	add := func(i int, sign int64) {
		if i == 0 || i == solver.n+1 {
			return
		}
		if a[i-1] < a[i] && a[i] > a[i+1] {
			solver.ans += sign * int64(a[i])
		}
		if a[i-1] > a[i] && a[i] < a[i+1] {
			solver.ans -= sign * int64(a[i])
		}
	}

	addTwo := func(sign int64) {
		if l+1 == r {
			add(l-1, sign)
			add(l, sign)
			add(r, sign)
			add(r+1, sign)
			return
		}
		add(l-1, sign)
		add(l, sign)
		add(l+1, sign)
		if l+2 != r {
			add(r-1, sign)
		}
		add(r, sign)
		add(r+1, sign)
	}

	addTwo(-1)
	solver.A[l], solver.A[r] = solver.A[r], solver.A[l]
	addTwo(1)

	return solver.ans
}
