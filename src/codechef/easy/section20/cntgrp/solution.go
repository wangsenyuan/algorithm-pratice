package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n-1)
		fmt.Println(solve(n, m, A))
	}
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

func pow(a, b int64) int64 {
	x, y := int64(1), a

	for b > 0 {
		if b&1 == 1 {
			x = (x * y) % MOD
		}
		y = (y * y) % MOD
		b >>= 1
	}
	return x
}

func solve(n int, m int, A []int) int {
	cnt := make(map[int]int)
	cnt[0] = 1
	for i := 0; i < len(A); i++ {
		x := A[i]
		cnt[x]++
	}
	if cnt[0] > 1 {
		return 0
	}
	var tot int64

	for _, v := range cnt {
		tot += int64(v) * int64(v-1) / 2
	}

	var prod int64 = 1

	for i := 0; i < len(A); i++ {
		prod *= int64(cnt[A[i]-1])
		prod %= MOD
	}

	ans := prod * nCr(tot, int64(m-n+1))
	ans %= MOD
	return int(ans)
}

func nCr(n int64, r int64) int64 {
	if n < r {
		return 0
	}
	if n == r || r == 0 {
		return 1
	}
	var numerator int64 = 1
	var denominator int64 = 1
	for i := int64(1); i <= r; i++ {
		denominator *= i
		denominator %= MOD
	}
	for i := int64(0); i < r; i++ {
		numerator *= n - i
		numerator %= MOD
	}
	ans := numerator * pow(denominator, MOD-2)
	ans %= MOD
	return ans
}
