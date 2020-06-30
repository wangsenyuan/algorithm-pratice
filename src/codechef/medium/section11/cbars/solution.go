package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	bs, _ := reader.ReadBytes('\n')

	var n int

	pos := readInt(bs, 0, &n)

	var m uint64

	readUint64(bs, pos+1, &m)

	fmt.Println(solve(n, m))
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

const MOD = 1000000007

func solve(n int, m uint64) int {
	N := 1 << uint(n)

	X := NewMatrix(N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			X.M[i][j] = int64(check(i, j, n))
		}
	}

	Y := Pow(X, m-1)

	var res int64

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			res = (res + Y.M[i][j]) % MOD
		}
	}
	return int(res)
}

func check(x, y int, n int) int {
	for i := 1; i < n; i++ {
		if x&(1<<uint(i-1)) > 0 && x&(1<<uint(i)) > 0 && y&(1<<uint(i-1)) > 0 && y&(1<<uint(i)) > 0 {
			return 0
		}
		if x&(1<<uint(i-1)) == 0 && x&(1<<uint(i)) == 0 && y&(1<<uint(i-1)) == 0 && y&(1<<uint(i)) == 0 {
			return 0
		}
	}

	return 1
}

type Mat struct {
	M [][]int64
	n int
}

func NewMatrix(n int) Mat {
	M := make([][]int64, n)
	for i := 0; i < n; i++ {
		M[i] = make([]int64, n)
	}

	return Mat{M, n}
}

func UnitMatrix(n int) Mat {
	m := NewMatrix(n)

	for i := 0; i < n; i++ {
		m.M[i][i] = 1
	}

	return m
}

func (this Mat) Mul(that Mat) Mat {
	res := NewMatrix(this.n)

	for i := 0; i < this.n; i++ {
		for j := 0; j < that.n; j++ {
			for k := 0; k < this.n; k++ {
				res.M[i][j] += (this.M[i][k] * that.M[k][j]) % MOD
				res.M[i][j] %= MOD
			}
		}
	}

	return res
}

func Pow(a Mat, b uint64) Mat {
	r := UnitMatrix(a.n)

	for b != 0 {
		if b&1 == 1 {
			r = r.Mul(a)
		}
		a = a.Mul(a)
		b >>= 1
	}

	return r
}
