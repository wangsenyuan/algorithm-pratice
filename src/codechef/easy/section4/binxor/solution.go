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

func readNum(scanner *bufio.Scanner) (a int) {
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		scanner.Scan()
		A := scanner.Text()
		scanner.Scan()
		B := scanner.Text()
		fmt.Println(solve(n, A, B))
	}
}

const MOD = 1000000007

func pow(a int, b int) int {
	A := int64(a)
	res := int64(1)

	for b > 0 {
		if b&1 == 1 {
			res *= A
			res %= MOD
		}

		A *= A
		A %= MOD
		b >>= 1
	}

	return int(res)
}

func modInverse(n int) int {
	return pow(n, MOD-2)
}

var fact []int

const N = 500000

func init() {
	fact = make([]int, N+1)

	fact[0] = 1
	fact[1] = 1

	for i := 2; i <= N; i++ {
		x := int64(i) * int64(fact[i-1])
		x %= MOD
		fact[i] = int(x)
	}
}

func nCr(n int, r int) int {
	res := int64(fact[n])
	res *= int64(modInverse(fact[r]))
	res %= MOD
	res *= int64(modInverse(fact[n-r]))
	res %= MOD
	return int(res)
}

func solve(n int, A, B string) int {
	var a, b int

	for i := 0; i < n; i++ {
		if A[i] == '1' {
			a++
		}
		if B[i] == '1' {
			b++
		}
	}

	x := abs(a - b)

	u := max(a, b)
	v := min(a, b)

	y := u + v

	if a+b >= n {
		y = 2*n - y
	}

	var ans int

	for i := x; i <= y; i += 2 {
		ans += nCr(n, i)
		if ans >= MOD {
			ans -= MOD
		}
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
