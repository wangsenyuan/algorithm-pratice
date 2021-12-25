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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		ss := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			ss[i] = scanner.Text()
		}
		fmt.Println(solve(n, ss))
	}
}

func solve(n int, ss []string) int64 {
	dp := make([]int64, 32)

	for i := 0; i < n; i++ {
		m := nomalize(ss[i])
		dp[m]++
	}
	res := dp[31] * (dp[31] - 1) / 2
	for i := 1; i < 32; i++ {
		for j := i + 1; j < 32; j++ {
			if i|j == 31 {
				res += dp[i] * dp[j]
			}
		}
	}

	return res
}

func nomalize(s string) int {
	var mask int
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			mask |= 1
		} else if s[i] == 'e' {
			mask |= 2
		} else if s[i] == 'i' {
			mask |= 4
		} else if s[i] == 'o' {
			mask |= 8
		} else {
			mask |= 16
		}
	}
	return mask
}
