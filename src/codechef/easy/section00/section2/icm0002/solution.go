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

const MOD = 998244353

const N = 1000010

var dp1 [N]int
var dp2 [N]int

func init() {
	var inv [N]int64
	inv[0] = 1
	inv[1] = 1
	for i := 2; i < N; i++ {
		inv[i] = inv[MOD%i] * int64(MOD-MOD/i) % MOD
	}

	var s1 [N]int64
	for i := 2; i < N; i++ {
		a := (s1[i-1] - s1[i/2] + int64((i-1)/2) + MOD) % MOD
		a = (a * 2) % MOD
		if i%2 == 0 {
			a = (a + int64(dp1[i/2]) + 1) % MOD
		}
		a = (a * inv[i-1]) % MOD
		dp1[i] = int(a)
		s1[i] = (s1[i-1] + a) % MOD
	}
	var s2 [N]int64

	for t := 2; t < N; t++ {
		a := (s2[t-1] - s2[t/2] + MOD) % MOD
		b := (s1[t-1] - s1[t/2] + MOD) % MOD
		b = (b * 2) % MOD
		a = (a + b + int64((t-1)/2)) % MOD
		a = (a * 2) % MOD
		if t%2 == 0 {
			a = (a + int64(dp2[t/2])) % MOD
			a = (a + 2*int64(dp1[t/2]) + 1) % MOD
		}
		a = (a * inv[t-1]) % MOD
		dp2[t] = int(a)
		s2[t] = (s2[t-1] + a) % MOD
	}
}

func solve(n int, m int) int {
	ans := int64(dp2[n]+dp2[m]) % MOD
	ans = (ans + 2*int64(dp1[n])*int64(dp1[m])) % MOD
	return int(ans)
}
