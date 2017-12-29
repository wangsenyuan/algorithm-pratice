package main

import (
	"bufio"
	"fmt"
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
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		scanner.Scan()
		n := readNum(scanner)
		grid := make([][]byte, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			grid[j] = scanner.Bytes()
		}
		fmt.Println(solve(n, grid))
	}
}

var dirs = []int{-1, 0, 1, 0, -1}

func solve(n int, grid [][]byte) int {
	stars := make([][]int, 15)
	sz := 1
	stars[0] = []int{0, 0}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '*' {
				stars[sz] = []int{i, j}
				sz++
			}
		}
	}
	stars[sz] = []int{n - 1, n - 1}
	// sz++

	ds := make([][]int, n)
	for i := 0; i < n; i++ {
		ds[i] = make([]int, n)
	}

	que := make([]int, n*n)
	calDist := func(from int, end int) int {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ds[i][j] = -1
			}
		}
		head, tail := 0, 0
		que[tail] = stars[from][0]*n + stars[from][1]
		tail++
		ds[stars[from][0]][stars[from][1]] = 0
		for head < tail {
			tmp := tail
			for head < tmp {
				v := que[head]
				a, b := v/n, v%n
				if a == stars[end][0] && b == stars[end][1] {
					return ds[a][b]
				}
				head++

				for k := 0; k < 4; k++ {
					c, d := a+dirs[k], b+dirs[k+1]
					if c >= 0 && c < n && d >= 0 && d < n && grid[c][d] != '#' && ds[c][d] == -1 {
						ds[c][d] = ds[a][b] + 1
						que[tail] = c*n + d
						tail++
					}
				}

			}
		}

		return -1
	}

	dist := make([][]int, sz+1)
	for i := 0; i <= sz; i++ {
		dist[i] = make([]int, sz+1)
	}

	for i := 0; i <= sz; i++ {
		// why this?
		dist[i][i] = 10000
		for j := i + 1; j <= sz; j++ {
			tmp := calDist(i, j)
			if tmp < 0 {
				return -1
			}
			dist[i][j] = tmp
			dist[j][i] = tmp
		}
	}
	sz--
	size := 1 << uint(sz)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, sz+1)
		for j := 0; j <= sz; j++ {
			dp[i][j] = 10000
		}
	}
	for i := 0; i < sz; i++ {
		dp[1<<uint(i)][i+1] = dist[0][i+1]
	}

	for i := 1; i < size; i++ {
		for j := 0; j < sz; j++ {
			if i&(1<<uint(j)) > 0 {
				// star j is used
				for k := 0; k < sz; k++ {
					if i&(1<<uint(k)) == 0 {
						//try use star k
						dp[i|(1<<uint(k))][k+1] = min(dp[i|(1<<uint(k))][k+1], dp[i][j+1]+dist[j+1][k+1])
					}
				}
			}
		}
	}
	best := 10000
	for i := 1; i <= sz; i++ {
		// take star i as the last stop before go to bottom-right
		best = min(best, dp[size-1][i]+dist[i][sz+1])
	}
	return best
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
