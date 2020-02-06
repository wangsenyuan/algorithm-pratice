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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		L := scanner.Text()
		scanner.Scan()
		R := scanner.Text()
		res := solve(str(L), str(R))
		fmt.Println(res)
	}
}

func str(s string) string {
	var i int

	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[i+1:]
}

const MAX_N = 100007
const MOD = 1000000007

var P []int64

func init() {
	P = make([]int64, MAX_N)
	P[0] = 1
	for i := 1; i < MAX_N; i++ {
		P[i] = (10 * P[i-1]) % MOD
	}
}

func solve(L, R string) int {
	_, b := F(R)
	a, _ := F(L)
	ans := b - a
	if ans < 0 {
		ans += MOD
	}
	return ans
}

func F(X string) (int, int) {
	num := make([][][]int64, 2)
	dp := make([][][]int64, 2)

	for i := 0; i < 2; i++ {
		num[i] = make([][]int64, 10)
		dp[i] = make([][]int64, 10)
		for j := 0; j < 10; j++ {
			num[i][j] = make([]int64, 2)
			dp[i][j] = make([]int64, 2)
		}
	}

	n := len(X)
	num[0][0][1] = 1

	for idx := 0; idx < n; idx++ {
		a := idx % 2
		b := 1 - a

		for digit := 0; digit < 10; digit++ {
			for eq := 0; eq < 2; eq++ {
				num[b][digit][eq] = 0
				dp[b][digit][eq] = 0
			}
		}

		for digit := 0; digit < 10; digit++ {
			for eq := 0; eq < 2; eq++ {
				for next := 0; next < 10; next++ {
					if eq == 1 && next > int(X[idx]-'0') {
						break
					}
					var neq int
					if eq == 1 && next == int(X[idx]-'0') {
						neq = 1
					}
					modAdd(&num[b][next][neq], num[a][digit][eq])
					modAdd(&dp[b][next][neq], dp[a][digit][eq])
					if next != digit {
						tmp := num[a][digit][eq] * P[n-1-idx]
						tmp %= MOD
						tmp *= int64(next)
						tmp %= MOD
						modAdd(&dp[b][next][neq], tmp)
					}
				}
			}
		}
	}

	var first, second int64

	p := n % 2

	for digit := 0; digit < 10; digit++ {
		modAdd(&first, dp[p][digit][0])
		modAdd(&second, dp[p][digit][0])
		modAdd(&second, dp[p][digit][1])
	}

	return int(first), int(second)
}

func modAdd(x *int64, y int64) {
	*x += y
	if *x >= MOD {
		*x -= MOD
	}
}
