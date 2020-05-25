package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		first := readNNums(reader, 4)
		N, m, x, y := first[0], first[1], first[2], first[3]
		A := readNNums(reader, N)
		fmt.Println(solve(N, m, x, y, A))
	}
}

const MOD = 1000000007

const MAX_M = 30

var coeff []int64

func init() {
	coeff = make([]int64, MAX_M+1)
	coeff[0] = 1
	for i := 1; i <= MAX_M; i++ {
		coeff[i] = (3*coeff[i-1] - 1) % MOD
	}
}

func solve(N, m, x, y int, A []int) int {

	F := func(m int, x int) int64 {
		var t int64
		var i int

		for 1<<uint(m) <= x {
			t += int64(A[i])*coeff[m] + int64(A[i+1])*(coeff[m]-1)
			t %= MOD
			i++
			x -= 1 << uint(m)
		}

		l := A[i]
		var r int
		if i+1 < N {
			r = A[i+1]
		}
		for m > 0 {
			m--
			if 1<<uint(m) <= x {
				t += int64(l)*coeff[m] + int64(l+r)*(coeff[m]-1)
				t %= MOD
				l += r
				x -= 1 << uint(m)
			} else {
				r += l
			}
		}

		return t
	}

	return int((F(m, y) - F(m, x-1) + MOD) % MOD)
}
