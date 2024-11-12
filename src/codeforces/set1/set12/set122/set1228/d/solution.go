package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) []int {
	res := make([]int, n+1)
	g := make([]map[int]int, n+1)

	for i := 1; i <= n; i++ {
		res[i] = 1
		g[i] = make(map[int]int)
	}

	for _, cur := range edges {
		u, v := cur[0], cur[1]
		g[u][v] = v
		g[v][u] = u
	}

	// 和1有边的，都不在集合1中
	for v := range g[1] {
		res[v] = 2
	}

	// 这些用来组成集合2和3
	if len(g[1]) < 2 {
		return nil
	}
	var second int
	for u := range g[1] {
		second = u
		break
	}
	// 那么和second有一条边，但是不在集合1中的，就在集合3中
	for u := range g[second] {
		if res[u] == 2 {
			res[u] = 3
		}
	}

	cnt := make([]int, 4)

	for u := 1; u <= n; u++ {
		cnt[res[u]]++
	}

	if cnt[3] == 0 {
		return nil
	}

	for u := 1; u <= n; u++ {
		if len(g[u]) != n-cnt[res[u]] {
			return nil
		}
		for v := range g[u] {
			if res[u] == res[v] {
				return nil
			}
		}
	}

	return res[1:]
}
