package main

import (
	"bufio"
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

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		fmt.Println(solve(n, k))
	}
}

const MOD = 1000000007

func solve(n int, k int) int {
	// find the fist time reach N
	N := int64(n)
	K := int64(k)

	X := (N * N) % MOD

	if K == 1 {
		return int(X)
	}

	// then back to 0
	X = (X + N) % MOD
	var KK int64
	if K%2 == 0 {
		KK = K - 1
	} else {
		KK = K - 2
	}

	Y := X
	Y += (2 * ((N * (KK / 2) % MOD) % MOD)) % MOD
	Y %= MOD
	Y += ((1 + KK/2) * (KK / 2)) % MOD

	if K%2 == 0 {
		Y += N
	} else {
		Y += N
		Y %= MOD
		Y += (K / 2) * 2
		Y %= MOD
	}

	return int(Y)
}
