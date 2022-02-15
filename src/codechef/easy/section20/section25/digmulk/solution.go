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

		n, k, m := readThreeNums(reader)
		s := readString(reader)

		res := solve(n, s, k, m)

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

func solve(n int, S string, k, m int) int {
	// 每个都是独立的，所以可以单个去计算
	mat := NewMatrix(10, 10)
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%d", int64(i)*int64(k))
		for _, c := range s {
			mat[int(c-'0')][i]++
		}
	}

	cnt := NewMatrix(10, 1)
	for i := 0; i < n; i++ {
		cnt[int(S[i]-'0')][0]++
	}

	ans := mat.pow(m).mul(cnt)
	var sum int64

	for _, row := range ans {
		for _, x := range row {
			sum += x
			sum %= MOD
		}
	}

	return int(sum)
}

type Matrix [][]int64

func (this Matrix) mul(that Matrix) Matrix {
	data := make([][]int64, len(this))
	for i := 0; i < len(data); i++ {
		data[i] = make([]int64, len(that[0]))
	}

	for i := 0; i < len(this); i++ {
		for j := 0; j < len(that[0]); j++ {
			for k := 0; k < len(this[i]); k++ {
				data[i][j] += this[i][k] * that[k][j] % MOD
				data[i][j] %= MOD
			}
		}
	}
	return Matrix(data)
}

func (mat Matrix) pow(n int) Matrix {
	res := NewIdentityMatrix(10)

	for n > 0 {
		if n&1 == 1 {
			res = res.mul(mat)
		}
		mat = mat.mul(mat)
		n >>= 1
	}
	return res
}

func NewMatrix(r int, c int) Matrix {
	data := make([][]int64, r)
	for i := 0; i < r; i++ {
		data[i] = make([]int64, c)
	}
	return Matrix(data)
}

func NewIdentityMatrix(r int) Matrix {
	mat := NewMatrix(r, r)
	for i := 0; i < r; i++ {
		mat[i][i] = 1
	}
	return mat
}
