package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		res := solve(n, k, A)
		fmt.Println(res)
	}
}

func solve(n int, k int, A []int) int64 {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2001)
	}

	for x := 1; x < 2001; x++ {
		for i := 0; i < n; i++ {
			dp[i+1][x] = dp[i][x]
			if A[i] < x {
				dp[i+1][x]++
			}
		}
	}

	check := func(l int, r int) bool {
		// check to see there exists a number x, that dp[x] * m < k & dp[x+1] * m >= k
		w := r - l
		// repeat m times
		m := (k + w - 1) / w
		if (dp[r][2000]-dp[l][2000])*m < k {
			return false
		}
		x := sort.Search(2000, func(i int) bool {
			return (dp[r][i]-dp[l][i])*m >= k
		})
		if x == 0 {
			return false
		}

		f := ((dp[r][x] - dp[l][x]) - (dp[r][x-1] - dp[l][x-1]))

		if f >= 2000 || f < 1 {
			return false
		}
		// f have to be in the range [l....r]
		return dp[r][f+1]-dp[l][f+1] > dp[r][f]-dp[l][f]
	}

	var res int64
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			if check(l, r+1) {
				res++
			}
		}
	}
	return res
}
