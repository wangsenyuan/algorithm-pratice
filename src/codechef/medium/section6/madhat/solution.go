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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007

const MAX_N = 100007

var F []int64
var IF []int64

func init() {
	F = make([]int64, MAX_N)

	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1]
		F[i] %= MOD
	}

	IF = make([]int64, MAX_N)
	IF[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = int64(i+1) * IF[i+1]
		IF[i] %= MOD
	}
}

func pow(a int64, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
			r %= MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF[r]
	res %= MOD
	res *= IF[n-r]
	res %= MOD
	return res
}

func solve(n int, A []int) int {

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + A[i]
	}

	sgt := make([]int, 4*n)

	var build func(i int, s, e int)

	build = func(i int, s, e int) {
		if s+1 == e {
			sgt[i] = s
			return
		}

		mid := (s + e) / 2
		build(2*i+1, s, mid)
		build(2*i+2, mid, e)
		if arr[sgt[2*i+1]] >= arr[sgt[2*i+2]] {
			sgt[i] = sgt[2*i+1]
		} else {
			sgt[i] = sgt[2*i+2]
		}
	}

	build(0, 0, n)

	var query func(i int, s, e int, l, r int) int

	query = func(i int, s, e int, l, r int) int {
		if l <= s && e <= r {
			return sgt[i]
		}
		mid := (s + e) / 2

		ret := -1
		if l < mid {
			ret = query(2*i+1, s, mid, l, r)
		}

		if mid < r {
			if ret < 0 {
				ret = query(2*i+2, mid, e, l, r)
			} else {
				tmp := query(2*i+2, mid, e, l, r)
				if arr[ret] < arr[tmp] {
					ret = tmp
				}
			}
		}
		return ret
	}

	var dfs func(i, j int) int64

	dfs = func(i, j int) int64 {
		if i == j {
			return 1
		}

		k := query(0, 0, n, i, j)

		if k+A[k]+1 > j {
			return 0
		}

		res := nCr(j-i-1, k-i)
		res *= dfs(i, k)
		res %= MOD
		res *= dfs(k+1, j)
		res %= MOD
		return res
	}

	res := dfs(0, n)

	return int(res)
}
