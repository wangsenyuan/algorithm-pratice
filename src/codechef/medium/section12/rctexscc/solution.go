package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	m, n, k := readThreeNums(reader)

	res := solve(m, n, k)

	fmt.Println(res)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func pow(a, b int) int {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func inv(x int) int {
	return pow(x, MOD-2)
}

func modMul(a int, b int) int {
	A := int64(a)
	B := int64(b)
	return int(A * B % MOD)
}

func modAdd(a int, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(m, n int, k int) int {
	if m == 1 {
		return solve1(n, k)
	}

	if n == 1 {
		return solve1(m, k)
	}

	// k * k * (2 * k - 1) + (n - 1) * (2 * k ** 3 - 3 * k ** 2 + 1) / k ** 3
	a := modMul(k, modMul(k, modAdd(modMul(2, k), MOD-1)))
	b := modMul(modAdd(n, MOD-1), modAdd(modAdd(modMul(2, modMul(k, modMul(k, k))), MOD-modMul(3, modMul(k, k))), 1))
	c := modMul(k, modMul(k, k))

	res := modMul(modAdd(a, b), inv(c))

	return res
}

func solve1(n, k int) int {
	// (k + (n - 1) / (k - 1)) / k
	res := modAdd(k, modMul(n-1, inv(k-1)))
	res = modMul(res, inv(k))
	return res
}
