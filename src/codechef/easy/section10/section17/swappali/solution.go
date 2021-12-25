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
		n := readNum(scanner)
		scanner.Scan()
		s := scanner.Text()
		res := solve(n, s)
		if res < 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			fmt.Printf("%d\n", res)
		}
	}
}

func solve(n int, s string) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		if s[0] == s[1] {
			return 0
		}
		return -1
	}

	if n == 3 {
		if s[0] == s[2] {
			return 0
		}
		if s[0] == s[1] || s[1] == s[2] {
			return 1
		}
		return -1
	}

	h := n / 2
	dp := make([]int, h+1)

	for i := 0; i <= h; i++ {
		dp[i] = -1
	}
	var i int
	if s[0] == s[n-1] {
		dp[0] = 0
		i = 1
	} else {
		if s[0] != s[n-2] || s[1] != s[n-1] {
			return -1
		}
		dp[0] = -1
		dp[1] = 1
		i = 2
	}

	for i < h {
		j := n - 1 - i
		if s[i] == s[j] {
			// we can choose to not swap
			dp[i] = dp[i-1]
		}
		// any way we can choose to swap
		if s[i] == s[j+1] && s[i-1] == s[j] {
			// safe to swap
			if dp[i-2] >= 0 && (dp[i] < 0 || dp[i] > dp[i-2]+1) {
				dp[i] = dp[i-2] + 1
			}
		}

		i++
	}

	if n&1 == 1 {
		// abcab
		if dp[h-1] >= 0 {
			return dp[h-1]
		}
		//abbca
		if dp[h-2] < 0 {
			return -1
		}
		// dp[h-2] is valid, but dp[h-1] is not
		if s[h-1] == s[h] || s[h] == s[h+1] {
			return dp[h-2] + 1
		}
		return -1
	}

	return dp[h-1]
}
