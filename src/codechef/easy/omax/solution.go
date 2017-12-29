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
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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

func readTwo(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	bs := scanner.Bytes()
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		m, n := readTwo(scanner)

		if m == 0 {
			break
		}
		mat := make([][]int, m)
		for i := 0; i < m; i++ {
			mat[i] = readNNums(scanner, n)
		}
		fmt.Println(solve(m, n, mat))
	}
}

func solve(m, n int, mat [][]int) int {
	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i][j] = mat[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	subMin := make([][][][]int, m)
	subs := make([][][][]int, m)
	for i := 0; i < m; i++ {
		subMin[i] = make([][][]int, n)
		subs[i] = make([][][]int, n)
		for j := 0; j < n; j++ {
			subMin[i][j] = make([][]int, m)
			subs[i][j] = make([][]int, m)
			for k := 0; k < m; k++ {
				subMin[i][j][k] = make([]int, n)
				subs[i][j][k] = make([]int, n)
				for x := 0; x < n; x++ {
					subMin[i][j][k][x] = math.MaxInt32
					subs[i][j][k][x] = math.MinInt32
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for a := 0; a <= i; a++ {
				c := i - a
				for b := 0; b <= j; b++ {
					d := j - b
					tmp := sum[i][j]
					if c > 0 {
						tmp -= sum[c-1][j]
					}

					if d > 0 {
						tmp -= sum[i][d-1]
					}
					if c > 0 && d > 0 {
						tmp += sum[c-1][d-1]
					}

					if tmp < subMin[c][d][i][j] {
						subMin[c][d][i][j] = tmp
					}
					subs[c][d][i][j] = tmp
					if c < i {
						if subMin[c+1][d][i][j] < subMin[c][d][i][j] {
							subMin[c][d][i][j] = subMin[c+1][d][i][j]
						}
						if subMin[c][d][i-1][j] < subMin[c][d][i][j] {
							subMin[c][d][i][j] = subMin[c][d][i-1][j]
						}
					}
					if d < j {
						if subMin[c][d+1][i][j] < subMin[c][d][i][j] {
							subMin[c][d][i][j] = subMin[c][d+1][i][j]
						}
						if subMin[c][d][i][j-1] < subMin[c][d][i][j] {
							subMin[c][d][i][j] = subMin[c][d][i][j-1]
						}
					}
				}
			}
		}
	}

	best := math.MinInt32

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for a := 0; a < i-1; a++ {
				for b := 0; b < j-1; b++ {
					tmp := subs[a][b][i][j] - subMin[a+1][b+1][i-1][j-1]
					if tmp > best {
						best = tmp
					}
				}
			}
		}
	}

	return best
}
