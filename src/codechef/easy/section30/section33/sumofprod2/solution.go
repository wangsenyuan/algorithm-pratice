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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const N = 100010
const MOD = 998244353

var F []int
var I []int

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

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = modMul(res, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return res
}

func init() {
	F = make([]int, N)
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I = make([]int, N)
	I[N-1] = pow(F[N-1], MOD-2)

	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}

func solve(A []int) int {
	n := len(A)

	//  任何包含0的字符串都没有贡献, 只有全1的字符串贡献为1
	// 对于确定的s， 长度为m的连续1的字符串的贡献为 m * (m + 1) / 2
	// 假设s中1的个数为m， 在全排列的情况下，
	// m个1连续的情况有 c(n - m + 1, 1) * p(m) * p(n - m)
	// m - 1 个1连续，1个单独的情况 这个方向是死胡同
	// 从 (m - 1)选择x个放在它的左边，(m - 1 - x)中选择y个放在它的右边, 它的贡献就是 (x + 1 + y) * C(m - 1, x) * C(m - 1 - x, y) * P(x) * P(y) * P()....
	// 也走不通，因为会出现重复计数
	// m个1， 在 m + 1 个位置中插入0, 将其分割为段，每段的贡献是知道的
	// P(M) * (M + 1) ** (n - m)
	var m int
	for i := 0; i < n; i++ {
		m += A[i]
	}

	var ans int

	for k := 1; k <= m; k++ {
		tmp := nCr(m, k)
		tmp = modMul(tmp, F[k])
		tmp = modMul(tmp, n-k+1)
		tmp = modMul(tmp, F[n-k])
		ans = modAdd(ans, tmp)
	}

	return ans
}
