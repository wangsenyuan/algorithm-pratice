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

	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n := readNum(scanner)
		parent := readNNums(scanner, n-1)
		club := readNNums(scanner, n)
		K := readNNums(scanner, n)
		res := solve(n, parent, club, K)
		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

const MOD = 1000000007

func solve(n int, parent []int, club []int, level []int) []int64 {
	children := make([][]int, n)

	for i := 1; i < n; i++ {
		p := parent[i-1]
		if children[p] == nil {
			children[p] = make([]int, 0, n)
		}
		children[p] = append(children[p], i)
	}

	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[club[i]]++
	}

	var dfs func(v int)

	result := make([]int64, n)
	dp := make([]map[int]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int64)
	}

	dfs = func(v int) {
		var s int64
		if level[v] > 0 {
			s = dp[club[v]][level[v]-1]
		}

		for _, w := range children[v] {
			dfs(w)
		}

		if level[v] >= count[club[v]] {
			result[v] = 0
		} else if level[v] == 0 {
			result[v] = 1
		} else {
			result[v] += (dp[club[v]][level[v]-1] - s + MOD) % MOD
			result[v] %= MOD
		}
		dp[club[v]][level[v]] += result[v]
		dp[club[v]][level[v]] %= MOD
	}

	dfs(0)

	return result
}
