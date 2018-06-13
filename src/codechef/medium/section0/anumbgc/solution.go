package main

import (
	"bufio"
	"os"
	"fmt"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64 = 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		var N int64
		scanner.Scan()
		readInt64(scanner.Bytes(), 0, &N)

		p, q := solve(N)
		fmt.Printf("%d/%d\n", p, q)

		t--
	}
}

func solve(N int64) (p int64, q int64) {
	nums := transform(N)
	l := len(nums)
	//dp[d][i][tight][started][found] = count of first i digits in (N) that has digits D
	dp := make([][][][]int64, l+1)

	for j := 0; j <= l; j++ {
		dp[j] = make([][][]int64, 2)
		for k := 0; k < 2; k++ {
			dp[j][k] = make([][]int64, 2)
			for u := 0; u < 2; u++ {
				dp[j][k][u] = make([]int64, 2)
				dp[j][k][u][0] = -1
				dp[j][k][u][1] = -1
			}
		}
	}
	var theDigit int

	var fn func(i int, tight int, started int, found int) int64
	fn = func(i int, tight int, started int, found int) int64 {
		if dp[i][tight][started][found] >= 0 {
			return dp[i][tight][started][found]
		}
		var res int64

		if i == l {
			if started > 0 && found > 0 {
				res = 1
			} else {
				res = 0
			}
		} else {

			for a := 0; a < 10; a++ {
				if tight == 1 && a > nums[i] {
					continue
				}
				newTight := 0
				if tight == 1 && a == nums[i] {
					newTight = 1
				}
				newStarted := 0
				if started == 1 || a > 0 {
					newStarted = 1
				}
				newFound := 0
				if found == 1 || (newStarted == 1 && a == theDigit) {
					newFound = 1
				}

				res += fn(i+1, newTight, newStarted, newFound)
			}
		}

		dp[i][tight][started][found] = res
		return res
	}

	for theDigit < 10 {
		for j := 0; j <= l; j++ {
			for k := 0; k < 2; k++ {
				for u := 0; u < 2; u++ {
					dp[j][k][u][0] = -1
					dp[j][k][u][1] = -1
				}
			}
		}
		p += fn(0, 1, 0, 0)
		theDigit++
	}
	q = N * 10

	g := gcd(p, q)

	return p / g, q / g
}

func transform(N int64) []int {
	res := make([]int, 0, 64)
	for N > 0 {
		res = append(res, int(N%10))
		N /= 10
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
