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
const M = 100010

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

var F [M]int
var I [M]int

func init() {
	F[0] = 1
	for i := 1; i < M; i++ {
		F[i] = int(int64(i) * int64(F[i-1]) % MOD)
	}
	I[M-1] = pow(F[M-1], MOD-2)
	for i := M - 2; i >= 0; i-- {
		I[i] = int(int64(i+1) * int64(I[i+1]) % MOD)
	}
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := int64(F[n])
	res *= int64(I[r])
	res %= MOD
	res *= int64(I[n-r])
	res %= MOD
	return int(res)
}

func solve(n, m int) int {
	//try F win
	// _FS_FS_... FS_
	var win int
	for i := 0; i <= n; i++ {
		win += nCr(m+1, i)
		win %= MOD
	}
	var lose int
	for i := 1; i <= n+1; i++ {
		lose += nCr(m+1, i)
		lose %= MOD
	}
	res := (win + lose) % MOD
	return int(res)
}
