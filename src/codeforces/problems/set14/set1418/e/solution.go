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
	D := readNNums(reader, n)

	solver := NewSolver(n, D)
	var buf bytes.Buffer
	for m > 0 {
		m--
		a, b := readTwoNums(reader)
		res := solver.solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 998244353

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}

func inverse(num int64) int64 {
	return pow(num, MOD-2)
}

func modDiv(a, b int64) int64 {
	return a * inverse(b) % MOD
}

func modMul(a, b int64) int64 {
	return a * b % MOD
}

type Solver struct {
	sums []int64
	arr  []int
}

func NewSolver(n int, A []int) Solver {
	sort.Ints(A)
	sums := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = (sums[i] + int64(A[i])) % MOD
	}

	return Solver{sums, A}
}

func (solver Solver) solve(a, b int) int {
	n := len(solver.arr)

	i := sort.SearchInts(solver.arr, b)

	//if i < n && solver.arr[i] == b {
	//	i--
	//}

	big := n - i

	if a > big || a == n {
		return 0
	}
	tmp := (solver.sums[n] - solver.sums[i] + MOD) % MOD
	// tmp * (1 - a/big)
	res := modMul(tmp, modDiv(int64(big-a), int64(big)))

	tmp = solver.sums[i]

	res += modMul(tmp, modDiv(int64(big+1-a), int64(big+1)))

	res %= MOD

	return int(res)
}
