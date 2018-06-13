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
		firstLine := readNNums(scanner, 2)
		n, q := firstLine[0], firstLine[1]
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		colors := readNNums(scanner, n)
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(scanner, 2)
		}
		res := solve(n, edges, colors, q, queries)
		for i := 0; i < len(res); i++ {
			if res[i] {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		t--
	}
}

func solve(n int, edges [][]int, colors []int, q int, queries [][]int) []bool {
	neighbors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		neighbors[a][b] = true
		neighbors[b][a] = true
	}

	// l[x][y] = minmum black nodes that in tree rooted at x, with size y
	// r[x][y] = maxmum black nodes that in tree tooted at x, with size y
	l, r := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		l[i] = make([]int, n+1)
		r[i] = make([]int, n+1)
	}
	size := make([]int, n)
	// min[i] = minmum black nodes with size i
	// max[i] = maxmum black nodes with size i
	mn, mx := make([]int, n+1), make([]int, n+1)

	for i := 0; i <= n; i++ {
		mn[i] = 1 << 30
		mx[i] = 0
	}
	var dfs func(v, p int)

	dfs = func(v, p int) {
		l[v][1] = colors[v]
		r[v][1] = colors[v]
		size[v] = 1
		for w := range neighbors[v] {
			if p != w {
				dfs(w, v)
				size[v] += size[w]
			}
		}

		for i := 2; i <= size[v]; i++ {
			l[v][i] = n
			r[v][i] = 0
		}
		size[v] = 1
		for w := range neighbors[v] {
			if p != w {
				for j := size[v]; j > 0; j-- {
					for k := size[w]; k > 0; k-- {
						l[v][j+k] = min(l[v][j+k], l[v][j]+l[w][k])
						r[v][j+k] = max(r[v][j+k], r[v][j]+r[w][k])
					}
				}
				size[v] += size[w]
			}
		}
		for i := 1; i <= size[v]; i++ {
			mn[i] = min(mn[i], l[v][i])
			mx[i] = max(mx[i], r[v][i])
		}
	}

	dfs(0, -1)

	ans := make([]bool, q)

	for i := 0; i < q; i++ {
		s, b := queries[i][0], queries[i][1]
		if mn[s] <= b && mx[s] >= b {
			ans[i] = true
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
