package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		nums := readNNums(reader, 4)
		res := solve(nums[0], nums[1], nums[2], nums[3])

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

const N = 500

func solve(n, m, l, r int) int {
	var ans int

	mat := make(Mat, N)
	res := make(Mat, N)
	for i := 0; i < N; i++ {
		mat[i] = make([]int, N)
		res[i] = make([]int, N)
	}

	for d := l; d <= r; d++ {
		pf := primeFactors(d)

		p := len(pf)
		mat.Reset()
		res.Reset()

		for mask := 0; mask < 1<<p; mask++ {
			for i := 1; i <= m; i++ {
				next := mask
				for j := 0; j < p; j++ {
					k := i
					var c int
					for k%pf[j].first == 0 {
						c++
						k /= pf[j].first
					}
					if c >= pf[j].second {
						next |= 1 << j
					}
				}
				mat[mask][next]++
			}
		}

		k := 1 << (p + 1)
		for i, c := 1<<p, 0; i < k; i, c = i+1, c+1 {
			mat[i][c]++
		}
		for i := 0; i < k; i++ {
			res[i][i] = 1
		}
		for i := 0; (1 << i) <= n; i++ {
			if n&(1<<i) > 0 {
				res.Mul(mat, res, k)
			}
			mat.Mul(mat, mat, k)
		}
		ans = modAdd(ans, res[0][(1<<p)-1])
	}

	return ans
}

type Mat [][]int

func (a Mat) Reset() {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = 0
		}
	}
}

func (a Mat) Mul(b Mat, c Mat, n int) {
	//n := len(a)
	d := make(Mat, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				d[i][j] = modAdd(d[i][j], modMul(a[i][k], b[k][j]))
			}
		}
	}
	for i := 0; i < n; i++ {
		copy(c[i], d[i])
	}
}

func primeFactors(num int) []Pair {
	var res []Pair
	for x := 2; x <= num/x; x++ {
		if num%x == 0 {
			var c int
			for num%x == 0 {
				c++
				num /= x
			}
			res = append(res, Pair{x, c})
		}
	}
	if num > 1 {
		res = append(res, Pair{num, 1})
	}
	return res
}

type Pair struct {
	first  int
	second int
}
