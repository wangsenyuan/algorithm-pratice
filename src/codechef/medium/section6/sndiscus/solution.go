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

const MAX_N = 50

var snakes [][]int

func init() {
	snakes = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		snakes[i] = make([]int, 4)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)
	for tc > 0 {
		tc--

		n := readNum(scanner)
		for i := 0; i < n; i++ {
			fillNNums(scanner, 4, snakes[i])
		}
		fmt.Println(solve(n, snakes))
	}
}

func solve(n int, snakes [][]int) int {

	cal2 := func(x, x1, x2 int) int {
		if x1 > x2 {
			x1, x2 = x2, x1
		}

		if x1 <= x && x <= x2 {
			return 0
		}
		return min(abs(x1-x), abs(x2-x))
	}

	cal := func(x, y int) int {
		var dist int

		for i := 0; i < n; i++ {
			x1, y1, x2, y2 := snakes[i][0], snakes[i][1], snakes[i][2], snakes[i][3]
			var tmp int
			if x1 == x2 {
				tmp = abs(x-x1) + cal2(y, y1, y2)
			} else {
				tmp = abs(y-y1) + cal2(x, x1, x2)
			}
			dist = max(dist, tmp)
		}

		return dist
	}

	ans := math.MaxInt32
	for i := 1; i <= 50; i++ {
		for j := 1; j <= 50; j++ {
			ans = min(ans, cal(i, j))
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
