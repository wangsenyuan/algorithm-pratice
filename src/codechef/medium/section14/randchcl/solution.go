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
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		W := readNNums(reader, n)
		res := solve(n, W)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

const MOD = 998244353
const MAX_N = 300005
const MAX_W = 500005

var INV [MAX_N]int64
var C [MAX_W]int64
var D [MAX_W]int64
var E [MAX_W]int64

func init() {
	INV[1] = 1
	for i := 2; i < MAX_N; i++ {
		INV[i] = MOD - MOD/int64(i)*INV[MOD%i]%MOD
	}
}

func solve(n int, W []int) int {
	for i := 0; i < n; i++ {
		C[W[i]]++
	}

	for i := 1; i < MAX_W; i++ {
		for j := i; j < MAX_W; j += i {
			D[i] += C[j]
		}
	}

	var ans int64

	for i := MAX_W - 1; i > 0; i-- {
		E[i] = D[i] * D[i]
		for j := 2 * i; j < MAX_W; j += i {
			E[i] -= E[j]
		}
		ans = (ans + E[i]*int64(i)) % MOD
	}

	for i := 1; i < MAX_W; i++ {
		for j := i; j < MAX_W; j += i {
			D[i] -= C[j]
		}
	}

	for i := 0; i < n; i++ {
		C[W[i]]--
	}

	ans = ans * int64(n+1) % MOD * int64(2*n+1) % MOD * INV[6] % MOD * INV[n] % MOD * INV[n] % MOD

	return int(ans)
}
