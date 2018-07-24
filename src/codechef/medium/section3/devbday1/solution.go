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
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		friend := readNNums(scanner, n)
		gift := readNNums(scanner, n)
		res := solve(n, gift, friend)
		fmt.Printf("%d\n", res)
	}
}

func solve(n int, gift []int, friend []int) int64 {
	// fmt.Printf("[debug]-- %d %v %v\n", n, gift, friend)
	for i := 0; i < n; i++ {
		friend[i]--
	}

	outs := make([][]int, n)
	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for i := 0; i < n; i++ {
		j := friend[i]
		outs[j] = append(outs[j], i)
	}

	vis := make([]bool, n)

	var dfs2 func(x int) int64

	dfs2 = func(x int) int64 {
		vis[x] = true
		ans := int64(gift[x])

		for _, v := range outs[x] {
			ans += dfs2(v)
		}
		return max(0, ans)
	}

	in := make([]int, n)

	bfs := func(x int) int64 {
		var ans int64
		ans += int64(gift[x])

		for j := 0; j < len(outs[x]); j++ {
			if in[outs[x][j]] < in[x] {
				ans += dfs2(outs[x][j])
			}
		}

		for i := friend[x]; i != x; i = friend[i] {
			ans += int64(gift[i])
			for j := 0; j < len(outs[i]); j++ {
				if in[outs[i][j]] < in[x] {
					ans += dfs2(outs[i][j])
				}

			}
		}
		return max(0, ans)
	}

	var dfs func(u int, time *int) int64
	dfs = func(u int, time *int) int64 {
		if vis[u] {
			//a loop
			return bfs(u)
		}
		*time++
		in[u] = *time
		vis[u] = true
		return dfs(friend[u], time)
	}
	var ans int64
	for i := 0; i < n; i++ {
		if !vis[i] {
			ans += dfs(i, new(int))
		}
	}
	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
