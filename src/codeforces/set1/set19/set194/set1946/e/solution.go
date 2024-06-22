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
		n, m1, m2 := readThreeNums(reader)
		pref := readNNums(reader, m1)
		suf := readNNums(reader, m2)
		res := solve(n, pref, suf)
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

const MOD = 1_000_000_000 + 7

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

const MAX_N = 2_000_00 + 10

var F [MAX_N]int
var I [MAX_N]int

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = modMul(i, F[i-1])
	}
	I[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}

func solve(n int, pref []int, suf []int) int {
	m1 := len(pref)
	m2 := len(suf)
	if pref[0] != 1 || suf[m2-1] != n || pref[m1-1] != suf[0] {
		return 0
	}
	// 选这么多作为前缀
	ans := nCr(n-1, suf[0]-1)

	f := func(arr []int) int {
		res := 1
		for i := len(arr) - 2; i >= 0; i-- {
			// 选出一部分放在arr[i]的后面
			r := arr[i+1] - arr[i] - 1
			// arr[i+1]前面共有 arr[i+1]个数，除了最大的和第二大的不能选，其他的都可以选
			n := arr[i+1] - 2
			res = modMul(res, modMul(nCr(n, r), F[r]))
		}
		return res
	}

	ans = modMul(ans, f(pref))

	reverse(suf)
	for i := 0; i < len(suf); i++ {
		suf[i] = n - suf[i] + 1
	}

	ans = modMul(ans, f(suf))

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
