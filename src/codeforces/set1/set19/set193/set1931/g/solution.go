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
		cnt := readNNums(reader, 4)
		res := solve(cnt)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

// 1000000000
const MOD = 998244353

const N = int(100 + 2*1e6)

var F []int
var I []int

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
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

func solve(cnt []int) int {
	if cnt[0]+cnt[1] == 0 {
		if cnt[2] == 0 || cnt[3] == 0 {
			return 1
		}
		return 0
	}

	if abs(cnt[0]-cnt[1]) > 1 {
		return 0
	}
	var res int
	if cnt[0] >= cnt[1] {
		res += calc(cnt[0], cnt[1]+1, cnt[2], cnt[3])
	}
	if cnt[1] >= cnt[0] {
		res += calc(cnt[0]+1, cnt[1], cnt[2], cnt[3])
	}

	return res % MOD
}

func abs(num int) int {
	return max(num, -num)
}

func calc(c1, c2, c3, c4 int) int {
	// 1, 2, 1, 2, 1, 2..1
	// 4可以在第一个1的前面或者所有2的后面
	// 相当于用c2 + 1 个搁板，把c4个数给分开
	// nCr(n + m - 1, m)
	// 3只能在1的后面
	ans := nCr(c2+c4-1, c4)
	// 相当于用c1-1个搁板，把c3个数给分开
	ans = modMul(ans, nCr(c1+c3-1, c3))

	return ans
}
