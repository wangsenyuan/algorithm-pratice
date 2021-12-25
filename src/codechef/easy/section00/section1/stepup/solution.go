package main

import (
	"bufio"
	"fmt"
	"os"
)

func toIntPair(buf []byte) (int, int) {
	i, a, b := 0, 0, 0

	for buf[i] != ' ' {
		a = a*10 + int(buf[i]-'0')
		i++
	}

	i++

	for i < len(buf) {
		b = b*10 + int(buf[i]-'0')
		i++
	}

	return a, b
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := toInt(scanner.Bytes())
	// fmt.Printf("echo %d\n", t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, m := toIntPair(scanner.Bytes())
		edges := make([][]int, m)

		// fmt.Printf("echo %d %d\n", n, m)

		for i := 0; i < m; i++ {
			scanner.Scan()
			a, b := toIntPair(scanner.Bytes())
			edges[i] = []int{a, b}
		}
		res, found := solve(n, m, edges)
		if !found {
			fmt.Println("IMPOSSIBLE")
		} else {
			fmt.Println(res)
		}
	}
}

func solve(n int, m int, edges [][]int) (int, bool) {
	adj := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[b][a] = true
	}

	// 0 means not processed, 1 means processed, -1 means processing
	status := make([]int, n+1)
	labels := make([]int, n+1)

	var dfs func(v int) (int, bool)

	dfs = func(v int) (int, bool) {
		if labels[v] > 0 {
			return labels[v], true
		}

		if len(adj[v]) == 0 {
			labels[v] = 1
			status[v] = 1
			return 1, true
		}
		status[v] = -1
		tmp := 0
		for w := range adj[v] {
			if status[w] == -1 {
				// a loop back
				return -1, false
			}
			wr, wt := dfs(w)
			if !wt {
				return -1, false
			}
			if wr > tmp {
				tmp = wr
			}
		}
		status[v] = 1
		labels[v] = tmp + 1
		return labels[v], true
	}

	ans := 0
	for i := 1; i <= n; i++ {
		tmp, can := dfs(i)
		if !can {
			return -1, false
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans, true
}
