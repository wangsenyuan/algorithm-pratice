package main

import (
	"bufio"
	"os"
	"fmt"
)

var ps []int

const N = 1000000

func init() {
	ps = make([]int, N+1)
	sieve := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if !sieve[i] {
			for j := i; j <= N; j += i {
				if ps[j] == 0 {
					ps[j] = i
				}
			}
		}
	}
}

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

	n := readNum(scanner)

	A := readNNums(scanner, n)

	fmt.Println(solve(n, A))
}

func solve(n int, A []int) int64 {
	var maxa int
	for i := 0; i < n; i++ {
		if A[i] > maxa {
			maxa = A[i]
		}
	}

	cnt := make([]int, maxa+1)

	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	dp := make([]int, maxa+1)

	var ans int64

	for i := 1; i <= maxa; i++ {
		if i == 1 {
			dp[i] = 1
		} else {
			if ps[i/ps[i]] != ps[i] {
				dp[i] = -1 * dp[i/ps[i]]
			}
		}
		if dp[i] != 0 {
			var total int64
			for j := i; j <= maxa; j += i {
				total += int64(cnt[j])
			}
			ans += total * (total - 1) * (total - 2) / 6 * int64(dp[i])
		}
	}

	return ans
}
