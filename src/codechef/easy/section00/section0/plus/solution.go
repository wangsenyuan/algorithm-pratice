package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		i++
		sign = -1
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
	t := readNum(scanner)

	for t > 0 {
		n, m := readTwoNums(scanner)
		G := make([][]int, n)

		for i := 0; i < n; i++ {
			G[i] = readNNums(scanner, m)
		}

		fmt.Println(solve(n, m, G))

		t--
	}
}

func solve(n, m int, G [][]int) int {
	N := max(n, m) + 1
	stack := make([][]int, N)

	for i := 0; i < N; i++ {
		stack[i] = make([]int, N)
	}
	ps := make([]int, N)
	sum := make([]int, N)

	left, right := make([][]int, n), make([][]int, n)

	for i := 0; i < n; i++ {
		left[i] = make([]int, m)
		right[i] = make([]int, m)
		// stack[0] = 0
	}

	for i := 0; i < n; i++ {
		sum[i] = 0
		stack[i][0] = 0
		ps[i] = 1
	}

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			sum[i] += G[i][j]
			left[i][j] = sum[i] - stack[i][0]
			for ps[i] > 0 && stack[i][ps[i]-1] >= sum[i] {
				ps[i]--
			}
			stack[i][ps[i]] = sum[i]
			ps[i]++
		}
	}

	for i := 0; i < n; i++ {
		sum[i] = 0
		stack[i][0] = 0
		ps[i] = 1
	}

	for j := m - 1; j >= 0; j-- {
		for i := 0; i < n; i++ {
			sum[i] += G[i][j]
			right[i][j] = sum[i] - stack[i][0]
			for ps[i] > 0 && stack[i][ps[i]-1] >= sum[i] {
				ps[i]--
			}
			stack[i][ps[i]] = sum[i]
			ps[i]++
		}
	}

	top, bottom := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		top[i] = make([]int, m)
		bottom[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		sum[i] = 0
		stack[i][0] = 0
		ps[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum[j] += G[i][j]
			top[i][j] = sum[j] - stack[j][0]
			for ps[j] > 0 && stack[j][ps[j]-1] >= sum[j] {
				ps[j]--
			}
			stack[j][ps[j]] = sum[j]
			ps[j]++
		}
	}

	for i := 0; i < m; i++ {
		sum[i] = 0
		stack[i][0] = 0
		ps[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			sum[j] += G[i][j]
			bottom[i][j] = sum[j] - stack[j][0]
			for ps[j] > 0 && stack[j][ps[j]-1] >= sum[j] {
				ps[j]--
			}
			stack[j][ps[j]] = sum[j]
			ps[j]++
		}
	}

	best := math.MinInt32

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			tmp := top[i-1][j] + right[i][j+1] + bottom[i+1][j] + left[i][j-1] + G[i][j]
			best = max(best, tmp)
		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
