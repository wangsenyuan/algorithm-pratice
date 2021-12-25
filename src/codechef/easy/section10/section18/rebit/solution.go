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
		L, _ := reader.ReadString('\n')
		res := solve(L[:len(L)-1])
		fmt.Printf("%d %d %d %d\n", res[0], res[1], res[2], res[3])
	}
}

const MOD = 998244353

func solve(L string) []int64 {
	// n := len(L)

	var pos int

	var loop func() []int64

	loop = func() []int64 {
		if L[pos] == '#' {
			pos++
			return getBase()
		}

		// L[pos] == '('
		pos++
		a := loop()
		op := L[pos]
		pos++
		b := loop()
		pos++

		return merge(a, op, b)
	}

	return loop()
}

func merge(a []int64, op byte, b []int64) []int64 {
	c := make([]int64, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			k := calc(i, j, op)
			c[k] += a[i] * b[j]
			c[k] %= MOD
		}
	}
	return c
}

func calc(a, b int, op byte) int {
	if op == '&' {
		if a == b {
			return a
		}
		if a == 0 || b == 0 {
			return 0
		}
		if a == 1 || b == 1 {
			return a + b - 1
		}
		return 0
	}
	if op == '|' {
		if a == b {
			return a
		}
		if a == 1 || b == 1 {
			return 1
		}
		if a == 0 || b == 0 {
			return a + b
		}
		return 1
	}
	return a ^ b
}

func getBase() []int64 {
	res := make([]int64, 4)
	res[0] = 748683265
	res[1] = 748683265
	res[2] = 748683265
	res[3] = 748683265
	return res
}
