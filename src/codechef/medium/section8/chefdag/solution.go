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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		A := make([][]int, k)
		for i := 0; i < k; i++ {
			A[i] = readNNums(scanner, n)
		}
		cnt, res := solve(n, k, A)
		// fmt.Println(cnt)
		f.WriteString(fmt.Sprintf("%d\n", cnt))
		for i := 0; i < n; i++ {
			if i < n-1 {
				f.WriteString(fmt.Sprintf("%d ", res[i]))
			} else {
				f.WriteString(fmt.Sprintf("%d\n", res[i]))
			}
		}
	}
}

func solve(n int, k int, A [][]int) (int, []int) {
	adj := make([][]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			adj[i][j] = true
		}
		adj[i][i] = false
	}

	for _, a := range A {
		for i := 0; i < n; i++ {
			x := a[i] - 1
			for j := 0; j < i; j++ {
				y := a[j] - 1
				adj[y][x] = false
			}
		}
	}

	var dfs func(u int) bool

	pair := make([]int, n)

	for i := 0; i < n; i++ {
		pair[i] = -1
	}

	vis := make([]bool, n)

	dfs = func(u int) bool {
		vis[u] = true
		for v := 0; v < n; v++ {
			if adj[u][v] && (pair[v] < 0 || !vis[pair[v]] && dfs(pair[v])) {
				pair[v] = u
				return true
			}
		}
		return false
	}
	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			vis[j] = false
		}
		if dfs(i) {
			cnt++
		}
	}

	for i := 0; i < n; i++ {
		pair[i]++
	}

	return n - cnt, pair
}
