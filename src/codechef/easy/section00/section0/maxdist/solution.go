package main

import (
	"bufio"
	"fmt"
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

func readOneNum(scanner *bufio.Scanner) (a int) {
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
	t := readOneNum(scanner)

	for i := 0; i < t; i++ {
		n := readOneNum(scanner)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		res := solve(n, edges)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, edges [][]int) bool {
	if n <= 2 {
		return false
	}
	adj := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		adj[a][b] = true
		adj[b][a] = true
	}

	computeCenters := func() []int {
		que := make([]int, n)
		parent := make([]int, n)
		dist := make([]int, n)
		bfs := func(x int) int {
			for i := 0; i < n; i++ {
				dist[i] = -1
			}
			dist[x] = 0
			head, rear := 0, 0
			que[rear] = x
			rear++
			for head < rear {
				tmp := rear
				for head < tmp {
					v := que[head]
					head++
					for w := range adj[v] {
						if dist[w] < 0 {
							dist[w] = dist[v] + 1
							parent[w] = v
							que[rear] = w
							rear++
						}
					}
				}
			}
			var maxDist int
			maxAt := x
			for i := 0; i < n; i++ {
				if dist[i] > maxDist {
					maxDist = dist[i]
					maxAt = i
				}
			}
			return maxAt
		}

		x := bfs(0)
		y := bfs(x)

		path := make([]int, 0, n)
		for y != x {
			path = append(path, y)
			y = parent[y]
		}

		path = append(path, x)
		diameter := len(path)
		mid := diameter / 2
		if diameter%2 == 1 {
			return []int{path[mid]}
		}
		return []int{path[mid-1], path[mid]}
	}

	centers := computeCenters()

	var height func(x, p int) int

	height = func(x, p int) int {
		var h int
		for w, v := range adj[x] {
			if !v {
				continue
			}
			if w == p {
				continue
			}
			tmp := height(w, x)
			if tmp > h {
				h = tmp
			}
		}
		return h + 1
	}

	countMaxHeight := func(c int) int {
		cnt, ht := 0, 0

		for x, v := range adj[c] {
			if !v {
				continue
			}
			tmp := height(x, c)
			if tmp > ht {
				ht, cnt = tmp, 1
			} else if tmp == ht {
				cnt++
			}
		}
		return cnt
	}

	if len(centers) == 1 {
		if countMaxHeight(centers[0]) == 2 {
			return true
		}
		return false
	}

	c1, c2 := centers[0], centers[1]
	adj[c1][c2] = false
	adj[c2][c1] = false

	if countMaxHeight(c1) == 1 || countMaxHeight(c2) == 1 {
		return true
	}
	return false
}
