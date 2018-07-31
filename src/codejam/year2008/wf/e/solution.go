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
	// scanner := bufio.NewScanner(os.Stdin)

	// tc := readNum(scanner)
	var tc int
	fmt.Scanf("%d", &tc)
	for x := 1; x <= tc; x++ {
		// n, m := readTwoNums(scanner)
		var n, m int
		fmt.Scanf("%d %d", &n, &m)

		G := make([][]byte, n)
		for i := 0; i < n; i++ {
			var s string
			fmt.Scanf("%s", &s)
			G[i] = []byte(s)
		}
		res := solve(n, m, G)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n, m int, G [][]byte) int {

	valid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == '?' {
				var t int
				for k := 0; k < 4; k++ {
					a, b := i+dd[k], j+dd[k+1]
					if valid(a, b) && G[a][b] == '#' {
						t++
					}
				}
				if t >= 2 {
					G[i][j] = '.'
				}
			}
		}
	}
	var idx int

	c := make([]int, n*m)
	pos := make([][]int, n*m)
	net := make([][]int, n*m)
	rev := make([][]int, n)
	for i := 0; i < n; i++ {
		rev[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == '?' {
				var t int
				for k := 0; k < 4; k++ {
					a, b := i+dd[k], j+dd[k+1]
					if valid(a, b) && G[a][b] == '#' {
						t++
					}
				}
				c[idx] = 2 - t
				pos[idx] = []int{i, j}
				rev[i][j] = idx
				net[idx] = make([]int, 0, 4)
				idx++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == '?' {
				u := rev[i][j]
				for k := 0; k < 4; k++ {
					a, b := i+dd[k], j+dd[k+1]
					if valid(a, b) && G[a][b] == '?' {
						net[u] = append(net[u], rev[a][b])
					}
				}
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == '#' {
				var t int
				for k := 0; k < 4; k++ {
					a, b := i+dd[k], j+dd[k+1]
					if valid(a, b) && G[a][b] == '#' {
						t++
					}
				}
				ans += 4 - t
			}
		}
	}

	var path func(u int, flag int) bool
	vis := make([]int, idx)

	pair := make([][]int, idx)
	for i := 0; i < idx; i++ {
		pair[i] = make([]int, 2)
		pair[i][0] = -1
		pair[i][1] = -1
	}

	path = func(u int, flag int) bool {
		if vis[u] == flag {
			return false
		}
		vis[u] = flag

		for i := 0; i < len(net[u]); i++ {
			v := net[u][i]
			if pair[v][0] == u || pair[v][1] == u {
				continue
			}
			if pair[v][0] < 0 {
				pair[v][0] = u
				return true
			}
			if c[v] == 2 && pair[v][1] < 0 {
				pair[v][1] = u
				return true
			}
			if path(pair[v][0], flag) {
				pair[v][0] = u
				return true
			}
			if c[v] == 2 && path(pair[v][1], flag) {
				pair[v][1] = u
				return true
			}
		}
		return false
	}

	var flag int
	for i := 0; i < idx; i++ {
		ans += 2 * c[i]
		if (pos[i][0]+pos[i][1])%2 == 0 {
			flag++
			for c[i] > 0 && path(i, flag) {
				ans -= 2
				c[i]--
				flag++
			}
		}
	}
	fmt.Fprintf(os.Stderr, "[debug]%d %d => %d\n", n, m, ans)
	for i := 0; i < n; i++ {
		fmt.Fprintf(os.Stderr, "[debug]%s\n", string(G[i]))
	}

	return ans
}
