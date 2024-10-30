package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(k, n, edges)

	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func createGraph(n int, edges [][]int) []map[int]int {
	g := make([]map[int]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make(map[int]int)
	}

	for _, ed := range edges {
		u, v := ed[0], ed[1]
		g[u][v]++
		g[v][u]++
	}
	return g
}

func solve(k int, n int, edges [][]int) []int {
	g := createGraph(n, edges)

	// 先形成一条至少k+1长的边
	// 这个肯定是可以的
	path := make([]int, n)
	vis := make([]bool, n+1)

	var dfs func(u int, d int) int

	dfs = func(u int, d int) int {
		path[d] = u
		vis[u] = true
		var cnt int
		for v := range g[u] {
			if vis[v] {
				cnt++
				continue
			}
			x := dfs(v, d+1)
			if x > 0 {
				return x
			}
		}
		if cnt == len(g[u]) {
			// 所有的都在path中
			return d
		}
		return -1
	}

	d := dfs(1, 0)
	// 肯定是true
	for i := 0; i < d; i++ {
		if g[path[d]][path[i]] > 0 {
			return path[i : d+1]
		}
	}
	return nil
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
