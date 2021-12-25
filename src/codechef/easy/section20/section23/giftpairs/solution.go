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
		res := solve(n, A)

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

const N = 1010

var F [N]int64
var I [N]int64

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}
	I[N-1] = inverse(F[N-1])
	for i := N - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}
}

func inverse(x int64) int64 {
	return pow(x, MOD-2)
}

func pow(a int64, b int64) int64 {
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * a) % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return R
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res = res * I[r] % MOD
	res = res * I[n-r] % MOD
	return res
}

func solve(n int, A []int) int {
	vis := make([]bool, n+1)
	for _, num := range A {
		vis[num] = true
	}

	var s1, s2 int

	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		if A[i-1] == 0 {
			s1++
		} else {
			s2++
		}
	}

	rem := s1 + s2

	if rem == 0 {
		return 0
	}
	ans := F[rem]
	for i := 1; i <= s1; i++ {
		cur := nCr(s1, i) * F[rem-i] % MOD
		if i&1 == 1 {
			ans = (ans + MOD - cur) % MOD
		} else {
			ans = (ans + cur) % MOD
		}
	}
	return int(ans)
}
