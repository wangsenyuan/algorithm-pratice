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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		scanner.Scan()
		s := scanner.Text()
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(scanner, 2)
		}
		res := solve(n, q, s, queries)

		for _, r := range res {
			if r {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}

func solve(n, q int, s string, queries [][]int) []bool {

	m := 1
	for m <= n {
		m <<= 1
	}

	dp := make([]bool, m<<1)

	for i := 0; i+2 < n; i++ {
		if s[i] == s[i+1] || s[i] == s[i+2] || s[i+1] == s[i+2] {
			dp[i+m] = true
		}
	}

	for i := m - 1; i > 0; i-- {
		dp[i] = dp[i<<1] || dp[(i<<1)|1]
	}

	var loop func(left, right, start, end, i int) bool

	loop = func(left, right, start, end, i int) bool {
		if left > right {
			return false
		}
		if left == start && right == end {
			return dp[i]
		}

		mid := (start + end) / 2
		if right <= mid {
			return loop(left, right, start, mid, i<<1)
		}
		if left > mid {
			return loop(left, right, mid+1, end, i<<1|1)
		}

		return loop(left, right, start, mid, i<<1) || loop(left, right, mid+1, end, i<<1|1)
	}

	res := make([]bool, q)
	for i := 0; i < q; i++ {
		query := queries[i]
		l, r := query[0], query[1]
		if r-l+1 >= 3 {
			res[i] = loop(l-1, r-3, 0, m-1, 1)
		}
	}
	return res
}
