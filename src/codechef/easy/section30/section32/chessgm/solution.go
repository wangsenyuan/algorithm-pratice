package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	d, k := readTwoNums(reader)

	res := solve(k, d)

	fmt.Println(res)
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
		if s[i] == '\n' {
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

const MOD = 1000000007

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

func solve(k int, d int) int {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return 0
	}

	adj := make([][][][]int, d+1)
	for i := 0; i <= d; i++ {
		adj[i] = make([][][]int, d+1)
		for j := 0; j <= d; j++ {
			adj[i][j] = make([][]int, d+1)
			for u := 0; u <= d; u++ {
				adj[i][j][u] = make([]int, d+1)
			}
		}
	}

	var states []Pair

	for i := 0; i <= d; i++ {
		for j := 0; j <= d; j++ {
			states = append(states, Pair{i, j})
			if i > 0 {
				adj[i][j][i-1][j] = 1
			}
			if i < d && j > 0 {
				adj[i][j][i+1][j-1] = 1
			}
			if j < d {
				adj[i][j][i][j+1] = 1
			}
		}
	}

	n := len(states)

	mat := NewMatrix(n, n)
	mat0 := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			(&mat).v[i][j] = adj[states[i].first][states[i].second][states[j].first][states[j].second]
			if i == 0 {
				(&mat0).v[i][j] = mat.v[i][j]
			}
		}
	}

	mat = pow(mat, k-1)
	mat0 = mat0.Mul(mat)

	var res int

	for i := 0; i < n; i++ {
		res = modAdd(res, mat0.v[i][0])
	}

	return res
}

type Pair struct {
	first  int
	second int
}
