package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	m, n := readTwoNums(reader)
	s := readNNums(reader, m)
	l := readNNums(reader, m)

	res := solve(n, s, l)

	fmt.Println(res)
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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

type Mat [][]int

func NewMat(n, m int) Mat {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	return res
}

func (this Mat) Mul(that Mat) Mat {
	n := len(this)
	res := NewMat(n, len(that[0]))

	for i := 0; i < n; i++ {
		for j := 0; j < len(that[0]); j++ {
			for k := 0; k < len(that); k++ {
				res[i][j] = add(res[i][j], mul(this[i][k], that[k][j]))
			}
		}
	}
	return res
}

func pow(m Mat, n int) Mat {
	res := NewMat(len(m), len(m))
	for i := 0; i < len(res); i++ {
		res[i][i] = 1
	}

	for n > 0 {
		if n&1 == 1 {
			res = res.Mul(m)
		}
		m = m.Mul(m)
		n >>= 1
	}
	return res
}

func solve(n int, s []int, l []int) int {
	m := len(s)

	mat := NewMat(m, m)

	t := make([]int, m)
	for i := 0; i < m; i++ {
		t[i] = add(s[i], l[i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			mat[i][j] = sub(mul(t[i], t[j]), mul(l[i], l[j]))
		}
	}

	res := pow(mat, n)

	v := NewMat(m, 1)
	v[0][0] = 1
	v = res.Mul(v)

	var ans int
	for i := 0; i < m; i++ {
		ans = add(ans, v[i][0])
	}

	return ans
}
