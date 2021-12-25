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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)

		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		res := solve(n, k, edges)
		fmt.Println(res)
	}
}

func solve(n, k int, edges [][]int) int {
	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}
	var res int
	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		var time int
		for _, v := range outs[u] {
			if p == v {
				continue
			}
			time = max(time, dfs(u, v))
		}

		time++
		if time <= k {
			res++
		}
		return time
	}

	dfs(-1, 0)

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
