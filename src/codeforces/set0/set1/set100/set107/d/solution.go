package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadBytes('\n')
	var n int64
	pos := readInt64(line, 0, &n)
	var m int
	readInt(line, pos+1, &m)
	labels := make([]string, m)
	S := make([]int, m)
	for i := 0; i < m; i++ {
		line, _ = reader.ReadBytes('\n')
		labels[i] = string(line[:1])
		readInt(line, 2, &S[i])
	}
	res := solve(n, labels, S)
	fmt.Println(res)
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && (bytes[i] >= '0' && bytes[i] <= '9') {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

const MOD = 12345

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(n int64, labels []string, S []int) int {
	m := len(labels)
	mul := make([]int, 26)
	for i := 0; i < 26; i++ {
		mul[i] = 1
	}
	mt := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		mt[i] = make([]bool, 124)
	}
	prod := 1

	for i := 0; i < m; i++ {
		lb := int(labels[i][0] - 'A')
		mul[lb] *= S[i]
		prod *= S[i]
		for j := 0; j < 123; j += S[i] {
			mt[lb][j] = true
		}
	}
	A := NewMat(1, prod)
	A.grid[0][0] = 1

	B := NewMat(prod, prod)

	hash := func(s []int) int {
		var res int
		for i := 0; i < 26; i++ {
			res = res*mul[i] + s[i]
		}
		return res
	}

	V := make([]bool, 125)
	F := make([]bool, 125)
	s := make([]int, 26)
	var dfs func() int
	dfs = func() int {
		x := hash(s)
		if V[x] {
			return x
		}
		V[x] = true
		F[x] = true
		for i := 0; i < 26; i++ {
			if s[i] > 0 && !mt[i][s[i]] {
				F[x] = false
			}
		}

		for i := 0; i < 26; i++ {
			if !mt[i][0] {
				continue
			}
			s[i] = (s[i] + 1) % mul[i]
			B.grid[x][dfs()]++
			s[i] = (s[i] + mul[i] - 1) % mul[i]
		}
		return x
	}

	dfs()

	A = A.Mul(B.Pow(n))
	var ans int

	for i := 0; i < prod; i++ {
		if F[i] {
			ans = modAdd(ans, A.grid[0][i]%MOD)
		}
	}
	return ans
}

type Mat struct {
	n    int
	m    int
	grid [][]int
}

func NewMat(n int, m int) Mat {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}
	return Mat{n, m, grid}
}

func (this Mat) Mul(that Mat) Mat {
	res := NewMat(this.n, that.m)
	for i := 0; i < this.n; i++ {
		for k := 0; k < this.m; k++ {
			if this.grid[i][k] != 0 {
				for j := 0; j < that.m; j++ {
					res.Add(i, j, modMul(this.grid[i][k], that.grid[k][j]))
				}
			}
		}
	}
	return res
}

func (this *Mat) Add(i, j int, v int) {
	this.grid[i][j] = modAdd(this.grid[i][j], v)
}

func (this Mat) Pow(n int64) Mat {
	res := NewMat(this.n, this.m)
	for i := 0; i < this.n; i++ {
		res.grid[i][i] = 1
	}
	a := this
	for n > 0 {
		if n&1 == 1 {
			res = res.Mul(a)
		}
		a = a.Mul(a)
		n >>= 1
	}
	return res
}
