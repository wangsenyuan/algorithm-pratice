package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * int64(sign)
	return i
}

func readNum(scanner *bufio.Scanner) (a int64) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	bs := scanner.Bytes()
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := int(readNum(scanner))
	ps := readNNums(scanner, n-1)
	parent := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		parent[i] = int(ps[i])
	}
	A := readNNums(scanner, n)
	res := solve(n, parent, A)

	fmt.Printf("%d", res[0])
	for i := 1; i < n; i++ {
		fmt.Printf(" %d", res[i])
	}
	fmt.Println()
}

func solve(n int, parent []int, A []int64) []int64 {
	children := make([]map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		children[i] = make(map[int]bool)
	}
	for i := 0; i < len(parent); i++ {
		p := parent[i]
		children[p][i+2] = true
	}
	result := make([]int64, n+1)
	stack := make([]int, n+1)
	depth := make([]int, n+1)
	var dfs func(v int, parent int)
	dfs = func(v int, parent int) {
		depth[v] = depth[parent] + 1
		stack[v] = parent
		for stack[v] > 0 && A[stack[v]-1] >= A[v-1] {
			stack[v] = stack[stack[v]]
		}
		result[v] = result[stack[v]] + A[v-1]*int64(depth[v]-depth[stack[v]])
		for w := range children[v] {
			dfs(w, v)
		}
	}

	dfs(1, 0)

	return result[1:]
}
