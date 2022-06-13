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
		N, L, R := readThreeNums(reader)
		res := solve(N, L, R)

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func pow(a int, b int) int {
	A := int64(a)
	res := int64(1)

	for b > 0 {
		if b&1 == 1 {
			res *= A
			res %= MOD
		}

		A *= A
		A %= MOD
		b >>= 1
	}

	return int(res)
}

func modInverse(n int) int {
	return pow(n, MOD-2)
}

var fact []int
var ifact []int

const N = 1000010

func init() {
	fact = make([]int, N+1)

	fact[0] = 1
	fact[1] = 1

	for i := 2; i <= N; i++ {
		x := int64(i) * int64(fact[i-1])
		x %= MOD
		fact[i] = int(x)
	}

	ifact = make([]int, N+1)
	ifact[N] = modInverse(fact[N])

	for i := N - 1; i >= 0; i-- {
		ifact[i] = int(int64(i+1) * int64(ifact[i+1]) % MOD)
	}

}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	res := int64(fact[n])
	res *= int64(ifact[r])
	res %= MOD
	res *= int64(ifact[n-r])
	res %= MOD
	return int(res)
}

func solve(N int, L int, R int) []int {
	// from N to 0, the result is nCr(n, 0), nCr(n - 1, 1)...
	m := R - L + 1
	res := make([]int, m)
	var p int
	for i := L; i <= R; i++ {
		x := (N + i) / 2
		if (N+i)%2 == 0 {
			res[p] = nCr(N, x)
		}
		p++
	}
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
