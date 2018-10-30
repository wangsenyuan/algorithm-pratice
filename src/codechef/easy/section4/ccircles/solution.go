package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

	N, Q := readTwoNums(scanner)
	points := make([][]int, N)
	for i := 0; i < N; i++ {
		points[i] = readNNums(scanner, 3)
	}
	queries := make([]int, Q)
	for i := 0; i < Q; i++ {
		queries[i] = readNum(scanner)
	}
	res := solve(N, Q, points, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

const MAX_K = 1e6 + 10

func solve(N, Q int, points [][]int, queries []int) []int {
	var res [MAX_K]int

	update := func(x, y float64) {
		xx := int(math.Ceil(x))
		yy := int(math.Floor(y))
		if xx < MAX_K {
			res[xx]++
		}
		if yy+1 < MAX_K {
			res[yy+1]--
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			x1, y1, r1 := points[i][0], points[i][1], points[i][2]
			x2, y2, r2 := points[j][0], points[j][1], points[j][2]

			curr1 := dist(x1, y1, x2, y2)
			curr2 := float64(r1 + r2)
			if curr1 > curr2 {
				min := curr1 - curr2
				max := curr1 + curr2
				update(min, max)
			} else {
				now := math.Abs(float64(r1-r2)) - curr1
				update(math.Max(now, 0), curr1+curr2)
			}
		}
	}

	for i := 1; i < MAX_K; i++ {
		res[i] += res[i-1]
	}

	ans := make([]int, Q)
	for i := 0; i < Q; i++ {
		ans[i] = res[queries[i]]
	}

	return ans
}

func dist(x1, y1, x2, y2 int) float64 {
	dx := int64(x1) - int64(x2)
	dy := int64(y1) - int64(y2)
	dd := float64(dx*dx + dy*dy)
	return math.Sqrt(dd)
}
