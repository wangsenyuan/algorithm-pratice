package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := readString(reader)
	b := readString(reader)
	k := readNum(reader)
	res := solve(k, a, b)
	fmt.Println(res)
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

const MOD = 1e9 + 7

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

func solve1(k int, start string, end string) int {
	n := len(start)
	// split 操作是把 x => ab => ba
	// n <= 1000
	// k <= 100000
	p := computeLps(end)

	dp := make([]int, k+1)
	fp := make([]int, k+1)
	var good int

	for i, j := 0, 0; i < 2*n; {
		if start[i%n] == end[j] {
			i++
			j++
		}
		if j == n {
			good++
			if i-n == 0 {
				dp[0]++
			}
			j = p[j-1]
		} else if i < 2*n && start[i%n] != end[j] {
			if j > 0 {
				j = p[j-1]
			} else {
				i++
			}
		}
	}
	good -= dp[0]

	if good == 0 {
		return 0
	}

	bad := n - good
	fp[0] = 1 - dp[0]
	for i := 1; i <= k; i++ {
		dp[i] = modAdd(modMul(dp[i-1], good-1), modMul(fp[i-1], good))
		fp[i] = modAdd(modMul(dp[i-1], bad), modMul(fp[i-1], bad-1))
	}

	return dp[k]
}

func computeLps(s string) []int {
	n := len(s)
	res := make([]int, n)

	for i := 1; i < n; i++ {
		j := res[i-1]
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		res[i] = j
		if s[i] == s[j] {
			res[i]++
		}
	}
	return res
}

func solve(k int, start string, end string) int {
	n := len(start)
	// split 操作是把 x => ab => ba
	// n <= 1000
	// k <= 100000
	p := computeLps(end)
	var first_good int
	var good int

	for i, j := 0, 0; i < 2*n; {
		if start[i%n] == end[j] {
			i++
			j++
		}
		if j == n {
			good++
			if i-n == 0 {
				first_good++
			}
			j = p[j-1]
		} else if i < 2*n && start[i%n] != end[j] {
			if j > 0 {
				j = p[j-1]
			} else {
				i++
			}
		}
	}
	good -= first_good

	if good == 0 {
		return 0
	}

	bad := n - good
	single := NewMat(2)
	single.Set(0, 0, good-1)
	single.Set(1, 0, good)
	single.Set(0, 1, bad)
	single.Set(1, 1, bad-1)

	res := single.Pow(k)
	dp := make([]int, 2)
	dp[0] = first_good
	dp[1] = 1 - first_good
	fp := make([]int, 2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			// new_dp[j] = (new_dp[j] + dp[i]*total.a[i][j]) % mod;
			fp[j] = modAdd(fp[j], modMul(dp[i], res.arr[i][j]))
		}
	}
	return fp[0]
}

type Mat struct {
	arr [][]int
}

func (this Mat) Mul(that Mat) Mat {
	res := make([][]int, len(this.arr))
	for i := 0; i < len(this.arr); i++ {
		res[i] = make([]int, len(that.arr[0]))
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			for k := 0; k < len(this.arr[0]); k++ {
				res[i][j] = modAdd(res[i][j], modMul(this.arr[i][k], that.arr[k][j]))
			}
		}
	}
	return Mat{res}
}

func (this Mat) Pow(n int) Mat {
	if n == 0 {
		return UnitMat(len(this.arr))
	}
	if n == 1 {
		return this
	}
	res := this.Pow(n / 2)
	res = res.Mul(res)
	if n&1 == 1 {
		return this.Mul(res)
	}
	return res
}

func NewMat(n int) Mat {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	return Mat{arr}
}

func (m *Mat) Set(i, j, v int) {
	m.arr[i][j] = v
}

func UnitMat(n int) Mat {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		arr[i][i] = 1
	}
	return Mat{arr}
}
