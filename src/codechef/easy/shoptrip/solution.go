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
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n, k := readTwo(scanner)
		coords := make([][]int, n)
		s := make([]string, n)
		for j := 0; j < n; j++ {
			a, b := readTwo(scanner)
			coords[j] = []int{a, b}
		}
		for j := 0; j < n; j++ {
			scanner.Scan()
			s[j] = scanner.Text()
		}
		fmt.Println(solve(n, coords, s, k))
	}
}

func solve(n int, coord [][]int, s []string, K int) float64 {
	goal := 1<<uint(K) - 1
	var flag int
	for i := 0; i < n; i++ {
		for j := 0; j < K; j++ {
			if s[i][j] == '1' {
				flag |= 1 << uint(j)
			}
		}
	}
	if flag != goal {
		return -1.0
	}

	dist := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]float64, n+1)
	}
	// number n is the cell (0, 0)
	for i := 0; i < n; i++ {
		a, b := coord[i][0], coord[i][1]
		for j := i + 1; j < n; j++ {
			c, d := coord[j][0], coord[j][1]
			dist[i][j] = distance(a, b, c, d)
			dist[j][i] = dist[i][j]
		}
		dist[i][i] = 0
		dist[n][i] = distance(0, 0, a, b)
		dist[i][n] = dist[n][i]
	}
	// f[state][i] = mim distance to get state from 0 to i
	f := make([][]float64, goal+1)

	for i := 0; i <= goal; i++ {
		f[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			f[i][j] = -1.0
		}
	}

	flags := make([]int, n)
	for i := 0; i < n; i++ {
		tmp := 0
		for j := 0; j < K; j++ {
			if s[i][j] == '1' {
				tmp |= 1 << uint(j)
			}
		}
		flags[i] = tmp
		f[tmp][i] = dist[i][n]
	}

	for state := 0; state <= goal; state++ {
		for i := 0; i < n; i++ {
			if f[state][i] < 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				tmp := state | flags[j]
				if f[tmp][j] == -1 || f[tmp][j] > f[state][i]+dist[i][j] {
					f[tmp][j] = f[state][i] + dist[i][j]
				}
			}
		}
	}

	best := math.MaxFloat64

	for i := 0; i < n; i++ {
		if f[goal][i] > 0 && f[goal][i]+dist[i][n] < best {
			best = f[goal][i] + dist[i][n]
		}
	}
	if best == math.MaxFloat64 {
		return -1.0
	}
	return best
}

func distance(a, b, c, d int) float64 {
	x := c - a
	y := d - b
	s := float64(x*x + y*y)
	return math.Sqrt(s)
}
