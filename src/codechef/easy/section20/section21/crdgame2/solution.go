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
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
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
const MAX_N = 100010

var F [MAX_N]int
var IF [MAX_N]int

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int((int64(i) * int64(F[i-1])) % MOD)
	}
	IF[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = int((int64(i+1) * int64(IF[i+1])) % MOD)
	}
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := int64(F[n])
	res *= int64(IF[r])
	res %= MOD
	res *= int64(IF[n-r])
	res %= MOD
	return int(res)
}

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

func solve(n int, A []int) int {
	var mx = A[0]
	var cnt = 1
	for i := 1; i < n; i++ {
		if A[i] == mx {
			cnt++
		} else if A[i] > mx {
			mx = A[i]
			cnt = 1
		}
	}
	x := pow(2, cnt)

	if cnt&1 == 0 {
		//to make it a draw, evenly distribute cnt
		y := nCr(cnt, cnt/2)
		x = (x - y + MOD) % MOD
	}
	z := pow(2, n-cnt)

	return int((int64(x) * int64(z)) % MOD)
}
