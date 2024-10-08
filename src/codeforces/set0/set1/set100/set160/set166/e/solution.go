package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	res := solve(n)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int) int {
	mat := NewMatrix(4, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mat.v[i][j] = 1
		}
		mat.v[i][i] = 0
	}

	mat = pow(mat, n)

	vec := NewMatrix(4, 1)
	vec.v[3][0] = 1

	res := mat.Mul(vec)

	return res.v[3][0]
}

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return a * b % MOD
}

type Mat struct {
	m int
	n int
	v [][]int
}

func NewMatrix(m int, n int) Mat {
	v := make([][]int, m)
	for i := 0; i < m; i++ {
		v[i] = make([]int, n)
	}
	return Mat{m, n, v}
}

func (this Mat) Mul(that Mat) Mat {
	m := this.m
	n := that.n
	v := make([][]int, m)
	for i := 0; i < m; i++ {
		v[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var sum int
			for k := 0; k < this.n; k++ {
				sum = modAdd(sum, modMul(this.v[i][k], that.v[k][j]))
			}
			v[i][j] = sum
		}
	}

	return Mat{m, n, v}
}

func pow(m Mat, n int) Mat {
	if n == 1 {
		return m
	}
	r := pow(m, n/2)
	r = r.Mul(r)
	if n%2 == 1 {
		r = r.Mul(m)
	}
	return r
}
