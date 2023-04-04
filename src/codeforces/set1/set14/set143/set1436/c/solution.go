package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x, p := readThreeNums(reader)

	res := solve(n, x, p)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

const N = 1010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}
	I[N-1] = inverse(F[N-1])

	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}

	return modMul(F[n], modMul(I[r], I[n-r]))
}

func solve(n, x, pos int) int {
	// n <= 1000
	// binary search, 假设 x在位置i, 怎么保证它一定会被找出来呢？
	//
	var dp, fp int
	// dp[i] = 找到i之前，二分查询需要查询几个数
	// A[pos] = x
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if pos < mid {
			// 要求 A[mid] < x
			dp++
			r = mid
		} else {
			// A[mid] > x
			fp++
			l = mid + 1
		}
	}

	// if A[i] = x
	a := modMul(nCr(n-x, dp), F[dp])
	// l 也有关系
	b := modMul(nCr(x-1, fp-1), F[fp-1])
	c := F[n-fp-dp]

	res := modMul(a, modMul(b, c))

	return res
}
