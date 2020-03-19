package main

import (
	"bufio"
	"fmt"
	"math"
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
		n, d := readTwoNums(scanner)
		fmt.Println(solve(n, d))
	}
}

func solve(n, d int) int64 {
	if d == 1 {
		return 0
	}

	deg := make([]int, n)
	deg[0] = 1
	for i := 1; i < n; i++ {
		deg[i] = deg[i-1] * d
	}

	numOfStates := d * deg[n-1]

	dp := make([]int64, numOfStates)

	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	dp[0] = 0

	coord := make([]int64, n)

	getState := func(state int) {
		for i := 0; i < n; i++ {
			coord[i] = int64(state % d)
			state /= d
		}
	}

	D := int64(d)
	for state := 0; state < numOfStates; state++ {
		getState(state)

		var sum int64
		var xor int64

		for i := 0; i < n; i++ {
			sum += coord[i]
			xor ^= coord[i]
		}

		for i := 0; i < n; i++ {
			if coord[i] < D-1 {
				dp[state+deg[i]] = min(dp[state+deg[i]], dp[state]+(sum+1)*(xor^coord[i]^(coord[i]+1)))
			}
		}
	}

	return dp[numOfStates-1]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
