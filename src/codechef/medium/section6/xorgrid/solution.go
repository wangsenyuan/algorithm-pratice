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

	n := readNum(scanner)

	grid := make([][]int, n)

	for i := 0; i < n; i++ {
		grid[i] = readNNums(scanner, n)
	}

	fmt.Println(solve(n, grid))
}

func solve(n int, grid [][]int) int {

	tries := make([]*Trie, n)

	var dfs func(i, j int, res int)

	dfs = func(i, j int, res int) {
		if i+j == n-1 {
			// the anti diagonal
			if tries[i] == nil {
				tries[i] = NewTrie(30)
			}
			tries[i].Add(res)
			return
		}

		dfs(i, j+1, res^grid[i][j+1])
		dfs(i+1, j, res^grid[i+1][j])
	}

	dfs(0, 0, grid[0][0])

	var ans int

	var dfs2 func(i, j int, res int)

	dfs2 = func(i, j int, res int) {
		if i-1+j == n-1 {
			// can go up/left to the anti diagonal
			ans = max(ans, max(tries[i-1].Xor(res), tries[i].Xor(res)))
			return
		}
		dfs2(i, j-1, res^grid[i][j-1])
		dfs2(i-1, j, res^grid[i-1][j])
	}

	dfs2(n-1, n-1, grid[n-1][n-1])

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Trie struct {
	children []*Trie
	pos      int
}

func NewTrie(pos int) *Trie {
	res := new(Trie)
	res.children = make([]*Trie, 2)
	res.pos = pos
	return res
}

func (trie *Trie) Add(num int) {
	if trie.pos < 0 {
		return
	}
	x := (num >> uint(trie.pos)) & 1

	if trie.children[x] == nil {
		trie.children[x] = NewTrie(trie.pos - 1)
	}

	trie.children[x].Add(num)
}

func (trie Trie) Xor(num int) int {
	if trie.pos < 0 {
		return 0
	}
	x := (num >> uint(trie.pos)) & 1
	y := x ^ 1
	if trie.children[y] != nil {
		ans := trie.children[y].Xor(num)
		return 1<<uint(trie.pos) | ans
	} else if trie.children[x] != nil {
		return trie.children[x].Xor(num)
	}
	return num & (1<<(uint(trie.pos)+1) - 1)
}
