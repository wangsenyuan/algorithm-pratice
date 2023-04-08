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
		n, m := readTwoNums(reader)
		res := solve(n, m)
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

func solve(N int, M int) int {
	// A[i] - mex(A[...i]) <= 1
	// 三种情况
	// A[i] = mex(A[...i])       (1)
	// A[i] = mex(A[...i]) + 1   (2)
	// A[i] = mex(A[...i]) - 1   (3)
	// n = 1, mex = 1, A[0] = 0, mex=1, M -1
	// n = 2, [0, 1], [1, 0], [M, 0]
	// mex[i] <= i
	// 如果全部是0, A[i] = 0, mex[i] = 1, 是满足条件的
	// 如果全部是1，mex[i] = 0, 也满足条件
	// 假设 A[i] = mex[i] - 1
	// 如果 A[i+1] = mex[i], 那么 mex[i+1] >= mex[i] + 1
	// 假设 mex[i+1] > mex[i] + 1, 需要存在某个j < i+1, A[j] = mex[i]+1
	//  j < i, 假设j = i - 1, A[i-1] = mex[i] + 1
	// 但是 A[i-1] <= mex[i-1] + 1 => mex[i] + 1 <= mex[i-1] + 1
	// mex[i] <= mex[i-1]
	// 但是 mex[i] >= mex[i-1], 所以要求mex[i] = mex[i-1]
	// 推广到一般的情况的, mex[j] = mex[j+1] ... = mex[i]
	// 一段连续的mex，A[j]的值 = mex + 1, 或者 mex - 1,
	//	直到遇到 A[i] = mex， 然后变成 mex + 2 (如果存在A[j] = mex+1), 或者 mex+1, (不存在 mex + 1)
	if M == 0 {
		return 1
	}
	if M == 1 {
		return N + 1
	}
	ans := 1

	for i := 0; i <= M; i++ {
		ans = modAdd(ans, nCr(N-1, i))
	}
	val := 1
	for i := 2; i <= N; i++ {
		// bad at i A[i] = mex + 1
		ans = modAdd(ans, modMul(val, pow(2, N-i)))
		val = modMul(val, 2)
		val = modSub(val, nCr(i-2, M-2))
	}

	return ans
}

const MOD = 1000000007

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
	return int(int64(a) * int64(b) % MOD)
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

const N = 1000010

var F [N]int
var I [N]int

func init() {

	F[0] = 1

	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I[N-1] = pow(F[N-1], MOD-2)

	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return modMul(F[n], modMul(I[r], I[n-r]))
}
