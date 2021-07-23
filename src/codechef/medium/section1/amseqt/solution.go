package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(n, m, A)
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

const MOD = 1000000009

func solve(n int, m int, A []int) int {
	M := 1 << uint(m)
	bit := make([]int, M+1)

	update := func(p int, v int) {
		p++
		for p <= M {
			bit[p] = modAdd(bit[p], v)
			p += p & -p
		}
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res = modAdd(res, bit[p])
			p -= p & -p
		}
		return res
	}

	query := func(l, r int) int {
		res := get(r) + MOD - get(l-1)
		return res % MOD
	}

	update(0, 1)

	var sum int

	M2 := M >> 1
	var res int
	for i := 0; i < n; i++ {
		sum += A[i]
		sum %= M
		if sum >= M2 {
			res = query(sum-M2+1, sum)
		} else {
			res = modAdd(query(0, sum), query(sum+M2+1, M-1))
		}
		update(sum, res)
	}

	return res
}

func modAdd(a, b int) int {
	a += b
	return a % MOD
}
