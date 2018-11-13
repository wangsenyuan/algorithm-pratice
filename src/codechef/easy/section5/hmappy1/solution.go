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
	line := readNNums(scanner, 3)
	n, k := line[0], line[1]
	A := readNNums(scanner, n)
	scanner.Scan()
	S := scanner.Text()
	res := solve(n, k, A, S)
	for _, num := range res {
		fmt.Println(num)
	}
}

func solve(n int, k int, A []int, S string) []int {
	dp := make([]int, n)
	lr := make([]int, n)
	rl := make([]int, n)
	pf := make([]int, n)
	sf := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = A[i]
		if i > 0 && A[i] > 0 {
			dp[i] += dp[i-1]
		}
		lr[i] = dp[i]
		if i > 0 && lr[i-1] > lr[i] {
			lr[i] = lr[i-1]
		}
		if dp[i] == i+1 {
			pf[i] = i + 1
		} else if i > 0 {
			pf[i] = pf[i-1]
		}
	}

	for i := n - 1; i >= 0; i-- {
		dp[i] = A[i]
		if i < n-1 && A[i] > 0 {
			dp[i] += dp[i+1]
		}
		rl[i] = dp[i]
		if i < n-1 && rl[i+1] > rl[i] {
			rl[i] = rl[i+1]
		}
		if dp[i] == n-i {
			sf[i] = n - i
		} else if i < n-1 {
			sf[i] = sf[i+1]
		}
	}

	res := make([]int, len(S))
	var p int
	j := n - 1
	for i := 0; i < len(S); i++ {
		if S[i] == '!' {
			j--
			if j < 0 {
				j = n - 1
			}
		} else {
			if j == n-1 {
				res[p] = min(lr[j], k)
				p++
			} else {
				jj := (j + 1) % n
				res[p] = max3(lr[j], rl[jj], pf[j]+sf[jj])
				res[p] = min(res[p], k)
				p++
			}
		}
	}
	return res[:p]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
