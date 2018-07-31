package main

import (
	"bufio"
	"fmt"
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

	line := readNNums(scanner, 3)
	n, m, k := line[0], line[1], line[2]
	G := make([][]int, m)
	for i := 0; i < m; i++ {
		G[i] = readNNums(scanner, 3)
	}
	res := solve(n, m, k, G)
	if res == nil {
		fmt.Println(-1)
	} else {
		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

func solve(n int, m int, k int, G [][]int) []int {
	net := make([][]int, n)
	ws := make([][]int, n)
	for i := 0; i < n; i++ {
		net[i] = make([]int, 0, (m+n)/n)
		ws[i] = make([]int, 0, (m+n)/n)
	}

	for i := 0; i < m; i++ {
		u, v, y := G[i][0]-1, G[i][1]-1, G[i][2]
		net[u] = append(net[u], v)
		ws[u] = append(ws[u], y)
		net[v] = append(net[v], u)
		ws[v] = append(ws[v], y)
	}
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = -1
	}
	que := make([]int, n)
	comp := make([]int, n)
	bfs := func(s int) bool {
		front, end := 0, 0
		que[end] = s
		end++
		x[s] = 0
		comp[s] = s
		for front < end {
			u := que[front]
			front++
			for i := 0; i < len(net[u]); i++ {
				v := net[u][i]
				y := ws[u][i]
				if x[v] >= 0 && x[v] != x[u]^y {
					return false
				}
				if x[v] < 0 {
					comp[v] = s
					x[v] = x[u] ^ y
					que[end] = v
					end++
				}
			}
		}
		return true
	}
	var c int
	for i := 0; i < n; i++ {
		if x[i] < 0 {
			if !bfs(i) {
				return nil
			}
			c = i
		}
	}

	for i := 0; i < n; i++ {
		if comp[i] == c {
			x[i] ^= (k - 1)
		}
	}

	return x
}
