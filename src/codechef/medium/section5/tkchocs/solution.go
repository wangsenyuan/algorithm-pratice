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
		N := readNum(scanner)
		grid := make([][]int, N)

		for i := 0; i < N; i++ {
			grid[i] = readNNums(scanner, i+1)
		}

		fmt.Println(solve(N, grid))
	}
}

func solve(N int, grid [][]int) int {
	po := make([][]int, N)
	mantis := make([][]int, N)

	for i := 0; i < N; i++ {
		po[i] = make([]int, N)
		mantis[i] = make([]int, N)
	}
	po[0][0] = grid[0][0]
	for i := 1; i < N; i++ {
		for j := 0; j <= i; j++ {
			var res int

			if j < i {
				res = max(res, po[i-1][j])
			}

			if j < i-1 {
				res = max(res, po[i-1][j+1])
			}
			if j > 0 {
				res = max(res, po[i-1][j-1])
			}

			po[i][j] = res + grid[i][j]
		}
	}

	mantis[N-1][N-1] = grid[N-1][N-1]

	for j := N - 2; j >= 0; j-- {
		for i := N - 1; i >= j; i-- {
			var res int
			if i > j {
				res = max(res, mantis[i][j+1])
			}
			if i > j+1 {
				res = max(res, mantis[i-1][j+1])
			}

			if i+1 < N {
				res = max(res, mantis[i+1][j+1])
			}
			mantis[i][j] = res + grid[i][j]
		}
	}

	var ans int

	for r, c, sum := N-1, 0, 0; r >= c; r, c = r-1, c+1 {
		if c+1 <= r {
			ans = max(ans, sum+po[r][c]+mantis[r][c+1])
		}

		if r+1 < N {
			ans = max(ans, sum+po[r][c]+mantis[r+1][c])
		}

		if c <= r-1 {
			ans = max(ans, sum+po[r-1][c]+mantis[r][c])
		}

		if c > 0 {
			ans = max(ans, sum+po[r][c-1]+mantis[r][c])
		}
		sum += grid[r][c]
		ans = max(ans, sum)
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
