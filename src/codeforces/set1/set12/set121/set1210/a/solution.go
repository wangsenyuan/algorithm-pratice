package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	// 可以给两个节点分配同一个数字吗？比如u, v都是2
	// 是可以的，只要给它们上面分配的边的总数，不超过6就可以了
	// 还是每个节点分配一个数字，但是同一个数字，它能够使用的指向边最多只有6
	marked := make([]bool, 36)

	check := func(p []int) int {
		clear(marked)
		var res int
		for _, e := range edges {
			u, v := p[e[0]-1], p[e[1]-1]
			u--
			v--
			u, v = min(u, v), max(u, v)
			val := u*6 + v
			if !marked[val] {
				res++
				marked[val] = true
			}
		}
		return res
	}

	a := make([]int, n)

	var res int
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = max(res, check(a))
			return
		}

		for j := 1; j <= 6; j++ {
			a[i] = j
			dfs(i + 1)
		}
	}

	dfs(0)

	return res
}
