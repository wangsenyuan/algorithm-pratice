package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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
func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const MOD = 1000000007

var L []int
var K []int

const N = 100010

func init() {
	// K[i] = no of 1 in F[i]
	// L[i] = len of F[i]
	K = make([]int, N)
	L = make([]int, N)
	// ans = K[i] % MOD * (pow(2, L[i], MOD - 1) % MOD
	K[0] = 0
	K[1] = 1
	L[0] = 1
	L[1] = 1

	for i := 2; i < N; i++ {
		K[i] = (K[i-2] + K[i-1]) % MOD
		L[i] = (L[i-2] + L[i-1]) % (MOD - 1)
	}
}

func pow(a, b, mod int) int {
	R := int64(1)
	A := int64(a)
	M := int64(mod)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % M
		}
		A = (A * A) % M
		b >>= 1
	}
	return int(R)
}

func solve(n int) int {
	res := int64(K[n]) * int64(pow(2, L[n]-1, MOD)) % MOD
	return int(res)
}
