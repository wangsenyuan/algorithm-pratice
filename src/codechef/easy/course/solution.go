package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readPair(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		scanner.Scan()
		r, R := readPair(scanner)
		n := readNum(scanner)
		cones := make([][]int, n)
		for j := 0; j < n; j++ {
			a, b := readPair(scanner)
			cones[j] = []int{a, b}
		}
		res := solve(r, R, n, cones)
		fmt.Printf("%.3f", res)
	}
}

func solve(r, R int, n int, cones [][]int) float64 {
	m := n + 2
	grid := make([][]float64, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			grid[i][j] = -1
		}
	}
	a, b := n, n+1
	for i := 0; i < n; i++ {
		x, y := cones[i][0], cones[i][1]
		z := math.Sqrt(float64(x*x + y*y))
		grid[a][i] = z - float64(r)
		grid[i][a] = grid[a][i]
		grid[b][i] = float64(R) - z
		grid[i][b] = grid[b][i]
		for j := i; j < n; j++ {
			if i == j {
				grid[i][j] = 0
			} else {
				u, v := cones[j][0], cones[j][1]
				z = math.Sqrt(float64((x-u)*(x-u) + (v-y)*(v-y)))
				grid[i][j] = z
				grid[j][i] = z
			}
		}
	}

	grid[a][b] = float64(R - r)
	grid[b][a] = grid[a][b]

	parent := make([]int, m)
	ks := make([]float64, m)
	for i := 0; i < m; i++ {
		ks[i] = math.MaxFloat64
	}
	set := make([]bool, m)

	ks[a] = 0
	parent[a] = -1

	for count := 0; count < m; count++ {
		u := -1
		for i := 0; i < m; i++ {
			if !set[i] && (u == -1 || ks[i] < ks[u]) {
				u = i
			}
		}
		set[u] = true
		for v := 0; v < m; v++ {
			if !set[v] && grid[u][v] < ks[v] {
				ks[v] = grid[u][v]
				parent[v] = u
			}
		}
	}
	var longestDistInThePath float64
	w := b
	for w != a {
		v := parent[w]
		if grid[v][w] > longestDistInThePath {
			longestDistInThePath = grid[v][w]
		}
		w = v
	}
	return longestDistInThePath
}
