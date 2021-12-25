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
		x, y := readTwoNums(reader)
		res := solve(x, y)
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

const M_X = 1000010

var F [M_X]int64
var I [M_X]int64

const MOD = 1000000007

func init() {
	F[0] = 1
	for i := 1; i < M_X; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}
	I[M_X-1] = pow(F[M_X-1], MOD-2)

	for i := M_X - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}
}

func pow(a int64, b int64) int64 {
	if b == 0 {
		return 1
	}
	x := pow(a, b/2)
	y := (x * x) % MOD
	if b%2 == 1 {
		return (y * int64(a)) % MOD
	}
	return y
}

func nCr(n, r int) int {
	if r < 0 || n < r {
		return 0
	}
	res := F[n]
	res *= I[r]
	res %= MOD
	res *= I[n-r]
	res %= MOD
	return int(res)
}

func findPos(num int) (int, int) {
	// num = r * (r - 1) / 2 + c
	N := int64(num)
	var left, right int64 = 0, M_X
	for left < right {
		mid := (left + right) / 2
		if mid*(mid-1)/2 >= N {
			right = mid
		} else {
			left = mid + 1
		}
	}
	r := int(right)
	r--
	x := r
	y := num - r*(r-1)/2
	return x, y
}

func solve(x int, y int) int {
	a, b := findPos(x)
	c, d := findPos(y)

	return nCr(c-a, d-b)
}
