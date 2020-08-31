package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readOneNum(scanner)

	for t > 0 {
		N, K := readTwoNums(scanner)
		scanner.Scan()
		S := scanner.Bytes()
		fmt.Println(solve(N, K, S))
		t--
	}
}

const MOD = 1000000007

const MAX = 300005

var p2 []int64
var p24 []int64
var p25 []int64
var fact []int64
var invp []int64
var f1 []int64

func init() {
	p2 = make([]int64, MAX)
	p24 = make([]int64, MAX)
	p25 = make([]int64, MAX)
	fact = make([]int64, MAX)
	invp = make([]int64, MAX)
	f1 = make([]int64, MAX)

	fact[0] = 1
	invp[0] = 1

	for i := 1; i < MAX; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}

	invp[MAX-1] = pow(fact[MAX-1], MOD-2)

	for i := int(MAX - 2); i >= 1; i-- {
		invp[i] = (invp[i+1] * int64(i+1)) % MOD
		if (fact[i]*invp[i])%MOD != 1 {
			panic("wrong")
		}
	}
	p2[0] = 1
	p24[0] = 1
	p25[0] = 1
	for i := 1; i < MAX; i++ {
		p2[i] = (p2[i-1] * 2) % MOD
		p24[i] = (p24[i-1] * 24) % MOD
		p25[i] = (p25[i-1] * 25) % MOD
	}
}

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

func ncr(n, r int) int64 {
	if n < r {
		return 0
	}
	ans := (fact[n] * invp[n-r]) % MOD
	ans = (ans * invp[r]) % MOD
	return ans
}

func solve(n, k int, S []byte) int {
	if k < 0 {
		return 0
	}

	fixed, nonFixed := 0, 0

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if S[i] == S[j] {
			fixed++
		} else {
			nonFixed++
		}
	}

	cal := func(k int) int64 {
		var ans int64
		f1[0] = 1

		for i := 1; i <= k; i++ {
			f1[i] = f1[i-1]
			if i%2 == 0 {
				f1[i] = (f1[i-1] + ncr(fixed, i/2)*p25[i/2]) % MOD
			}
		}

		for x := 0; x <= k; x++ {
			r := x + 2*(nonFixed-x)
			if r <= k && nonFixed >= x {
				f2 := (ncr(nonFixed, x) * p2[x]) % MOD
				f2 = (f2 * p24[nonFixed-x]) % MOD
				ans = (ans + f2*f1[k-r]) % MOD
			}
		}
		return ans
	}

	ans := cal(k)

	if n&1 == 1 {
		ans = (ans + (25 * cal(k-1))) % MOD
	}

	return int(ans)
}
